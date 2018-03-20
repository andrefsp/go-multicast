package gomulticast

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
)

type Discovery interface {
	GetNodes() []*Node
	Start()
}

// Constructor
func NewNodeDiscovery(clusterName, multicastAddress string) Discovery {
	if multicastAddress == "" {
		multicastAddress = defaultMulticastAddress
	}
	return &NodeDiscovery{
		clusterName:      clusterName,
		multicastAddress: multicastAddress,
	}
}

// Discovery message
type DiscoverMessage struct {
	ClusterName string `json:"cluster_name"`
	Message     string `json:"message"`
	Hostname    string `json:"hostname"`
}

// NodeDiscovery implements Discover interface
type NodeDiscovery struct {
	clusterName string

	// multicast address
	multicastAddress string

	// nodes
	nodes sync.Map
}

func (d *NodeDiscovery) GetNodes() []*Node {
	nodes := []*Node{}

	d.nodes.Range(func(key, value interface{}) bool {
		nodes = append(nodes, value.(*Node))
		return true
	})

	return nodes
}

func (d *NodeDiscovery) addOrUpdateNode(node *Node) {

	if _, ok := d.nodes.Load(node.Hostname); !ok {
		log.Printf("'%s' joined the cluster \n", node.Hostname)
	}

	d.nodes.Store(node.Hostname, node)
}

func (d *NodeDiscovery) checkNodes() {

	for range time.NewTicker(5 * time.Second).C {
		deletedNodes := []string{}
		// Collect nodes to delete
		d.nodes.Range(func(key, value interface{}) bool {
			node := value.(*Node)
			if node.LastHeartbeat.Add(time.Second * 5).Before(time.Now()) {
				deletedNodes = append(deletedNodes, node.Hostname)
			}
			return true
		})

		// delete them
		for _, hostname := range deletedNodes {
			log.Printf("'%s' abandoned the cluster \n", hostname)
			d.nodes.Delete(hostname)
		}
	}
}

func (d *NodeDiscovery) startListener() {
	Listen(d.multicastAddress, d.msgHandler)
}

func (d *NodeDiscovery) msgHandler(src *net.UDPAddr, n int, b []byte) {
	var message DiscoverMessage

	if err := json.Unmarshal(b[:n], &message); err != nil {
		log.Println("ERR:: ", err)
		return
	}

	if message.ClusterName == d.ClusterName {
		d.addOrUpdateNode(&Node{
			Hostname:      message.Hostname,
			IP:            strings.Split(src.String(), ":")[0],
			LastHeartbeat: time.Now(),
		})
	}
}

func (d *NodeDiscovery) startPing() {

	conn, err := NewBroadcaster(d.multicastAddress)
	if err != nil {
		log.Fatal(err)
	}

	for range time.NewTicker(2 * time.Second).C {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
		}

		data, err := json.Marshal(DiscoverMessage{
			ClusterName: d.clusterName,
			Message:     "Hello world",
			Hostname:    hostname,
		})
		if err != nil {
			log.Println(err)
		}

		conn.Write(data)
	}
}

func (d *NodeDiscovery) Start() {

	// Start ping
	go d.startPing()

	// Start listening
	go d.startListener()

	// Keep checking which nodes left the cluster
	go d.checkNodes()

	// Block here
	<-make(chan bool)
}
