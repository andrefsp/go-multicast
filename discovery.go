package gomulticast

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
)

var Nodes = map[string]*Node{}

type DiscoverMessage struct {
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
}

func CheckNodes() {
	for {
		time.Sleep(1 * time.Second)
		for hostname, node := range Nodes {
			if node.LastHeartbeat.Add(time.Second * 3).Before(time.Now()) {
				log.Printf("\t \t \t%s left the cluster \n ", hostname)
				delete(Nodes, hostname)
			}
		}
	}
}

func StartListener() {
	address := defaultMulticastAddress
	fmt.Printf("Listening on %s\n", address)
	Listen(address, msgHandler)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {

	var message struct {
		Message  string `json:"message"`
		Hostname string `json:"hostname"`
	}

	if err := json.Unmarshal(b[:n], &message); err != nil {
		log.Println("ERR:: ", err)
		return
	}

	_, ok := Nodes[message.Hostname]
	if !ok {
		fmt.Printf("'%s' joined the cluster\n", message.Hostname)
	}

	Nodes[message.Hostname] = &Node{
		Hostname:      message.Hostname,
		IP:            strings.Split(src.String(), ":")[0],
		LastHeartbeat: time.Now(),
	}
	fmt.Printf("%+v \n", Nodes)
}

func StartPinger() {
	address := defaultMulticastAddress
	ping(address)
}

func ping(addr string) {

	conn, err := NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
		}

		data, err := json.Marshal(DiscoverMessage{
			Message:  "Hello world",
			Hostname: hostname,
		})
		if err != nil {
			log.Println(err)
		}

		conn.Write(data)
		time.Sleep(3 * time.Second)
	}
}

/*
func main() {

	go func() {
		StartPinger()
	}()

	go func() {
		StartListener()
	}()

	go func() {
		CheckNodes()
	}()

	complete := make(chan bool)

	<-complete
}
*/
