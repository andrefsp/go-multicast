package gomulticast

import "time"

type Node struct {
	Hostname      string
	IP            string
	LastHeartbeat time.Time
}
