package main

import (
	"fmt"
	"time"

	gomulticast "github.com/andrefsp/go-multicast"
)

func main() {

	nodeDiscovery := gomulticast.NewNodeDiscovery("example-cluster", "")
	go nodeDiscovery.Start()

	for range time.NewTicker(5 * time.Second).C {
		fmt.Println("Current nodes in the cluster::: ")

		for _, node := range nodeDiscovery.GetNodes() {
			fmt.Printf("\t \t %+v \n", node)
		}

	}
	<-make(chan struct{})
}
