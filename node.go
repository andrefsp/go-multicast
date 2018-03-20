package gomulticast

import "time"

type Node struct {
	Hostname      string
	IP            string
	LastHeartbeat time.Time
}

func (n Node) GetHostname() string {
	return n.Hostname
}

func (n Node) GetIP() string {
	return n.IP
}
