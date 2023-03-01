package zmq

import (
	"fmt"
	"sync"

	zmq "github.com/pebbe/zmq4"
)

func InitializeZMQ(wg *sync.WaitGroup) {
	// ZMQ
	zctx, _ := zmq.NewContext()
	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.REQ)
	s.Connect("tcp://192.168.0.40:5555")
	wg.Add(1)
	go listenToResponse(s, wg)
	fmt.Println("ZMQ Initialized")
	wg.Done()
}

func MakeZMQRequest(s *zmq.Socket, message string) {
	s.Send(message, 0)
}

func listenToResponse(s *zmq.Socket, wg *sync.WaitGroup) {
	msg, _ := s.Recv(0)
	fmt.Println("RESPONSE: ", msg)
	wg.Done()
}

func createSubscriber() {
	zctx, _ := zmq.NewContext()
	sub, _ := zctx.NewSocket(zmq.SUB)
	sub.Connect("tcp://192.168.0.40:5556")
	sub.SetSubscribe((""))
	fmt.Println("\nStarting Listen:")
	for {
		msg, _ := sub.Recv(0)
		fmt.Println(msg)
		if msg == "stop" {
			break
		}
	}

}
