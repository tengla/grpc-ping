package ping

import (
	"encoding/json"
	"time"

	"github.com/tengla/grpc-ping/pong"
)

type pingServer struct {
	UnimplementedPingServiceServer
}

func NewPingServer() *pingServer {
	return &pingServer{}
}

func (p *pingServer) Ping(req *Empty, stream PingService_PingServer) error {
	for {
		pong := pong.NewPong("Keep me alive")
		data, err := json.Marshal(pong)
		if err != nil {
			return err
		}
		r := &PingResponse{
			Payload: string(data),
		}
		time.Sleep(time.Second)
		err = stream.Send(r)
		if err != nil {
			return err
		}
	}
}
