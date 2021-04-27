package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/tengla/grpc-ping/pong"
	ping "github.com/tengla/grpc-ping/protos"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	client := ping.NewPingServiceClient(conn)
	stream, err := client.Ping(context.Background(), &ping.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Ping(_) = _, %v", client, err)
		}
		var p pong.Pong
		err = json.Unmarshal([]byte(data.Payload), &p)
		if err != nil {
			log.Fatalf("json unmarshal error: %v", err)
		}
		fmt.Printf("%+v\n", p)
	}
}
