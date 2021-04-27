package ping

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/tengla/grpc-ping/pong"
)

type pingServer struct {
	UnimplementedPingServiceServer
}

func NewPingServer() *pingServer {
	return &pingServer{}
}

func pingChan() chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer func() {
			close(ch)
		}()
		messages := []string{"ABC", "DEF", "GHI", "JKL", "MNO", "PQR", "STU", "VWX", "YZ"}
		for {
			message := messages[int(rand.Int()%len(messages))]
			data, err := json.Marshal(
				pong.NewPong(message))
			if err != nil {
				ch <- err
			}
			r := &PingResponse{
				Payload: string(data),
			}
			ch <- r
			time.Sleep(time.Second * time.Duration(rand.Float32()*4))
		}
	}()
	return ch
}

// Ping - grpc stream
func (p *pingServer) Ping(req *Empty, stream PingService_PingServer) error {
	for v := range pingChan() {
		switch t := v.(type) {
		case *PingResponse:
			err := stream.Send(t)
			if err != nil {
				return err
			}
		case error:
			return t
		}
	}
	return nil
}
