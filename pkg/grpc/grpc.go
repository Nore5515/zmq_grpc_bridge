package grpc

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "bbb_bridge/pkg/grpc/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func InitializeGRPC(wg *sync.WaitGroup) {
	// GRPCS
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterGreeterServer(serv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// listen
	fmt.Println("GRPC Initalized")
	wg.Done()
}

func listenGRPC(serv *grpc.Server, lis net.Listener) {
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SendClientRequest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println("Client Request Sent!")
	return &pb.Response{Data: "OK"}, nil
}
