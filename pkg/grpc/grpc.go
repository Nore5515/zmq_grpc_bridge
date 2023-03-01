/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

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
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterGreeterServer(serv, &server{})

	// Start Listener
	wg.Add(1)
	go listenGRPC(serv, lis, wg)
	log.Printf("server listening at %v", lis.Addr())

	// Cleanup
	fmt.Println("GRPC Initalized")
	wg.Done()
}

func listenGRPC(serv *grpc.Server, lis net.Listener, wg *sync.WaitGroup) {
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	wg.Done()
}

func (s *server) SendClientRequest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println("Client Request Sent!")
	return &pb.Response{Data: "OK"}, nil
}
