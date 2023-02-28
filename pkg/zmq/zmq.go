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
	fmt.Println("ZMQ Initialized")
	wg.Done()
}

func makeZMQRequest(s *zmq.Socket, message string) {
	s.Send(message, 0)
}

func listenToResponse(s *zmq.Socket) {
	msg, _ := s.Recv(0)
	fmt.Println("RESPONSE: ", msg)
}

func createSubscriber(wg *sync.WaitGroup) {
	wg.Add(1)
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
	wg.Done()

}
