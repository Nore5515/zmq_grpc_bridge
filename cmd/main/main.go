
package main

import (
	"sync"

	grpc "bbb_bridge/pkg/grpc"
	zmq "bbb_bridge/pkg/zmq"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	// Initialization
	go grpc.InitializeGRPC(&wg)
	go zmq.InitializeZMQ(&wg)

	wg.Wait()
}
