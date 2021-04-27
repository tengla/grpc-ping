package main

import (
	"fmt"
	"log"
	"net"

	ping "github.com/tengla/grpc-ping/protos"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.ServerOption
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))
	if err != nil {
		log.Fatal(err.Error())
	}
	grpcServer := grpc.NewServer(opts...)
	ping.RegisterPingServiceServer(grpcServer, ping.NewPingServer())
	log.Println("Starting server at :9090")
	grpcServer.Serve(lis)
}
