package main

import gomulticast "github.com/dmichael/go-multicast"

func main() {

	go func() {
		gomulticast.StartPinger()
	}()

	go func() {
		gomulticast.StartListener()
	}()

	go func() {
		gomulticast.CheckNodes()
	}()

	complete := make(chan bool)

	<-complete
}
