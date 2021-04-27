package pong

import "time"

// Pong is payload sent to client
type Pong struct {
	Seq         int
	CurrentTime time.Time
	Message     string
}

var seq int

// NewPong creates a new Pong
func NewPong(msg string) Pong {
	seq += 1
	return Pong{
		Seq:         seq,
		CurrentTime: time.Now().UTC(),
		Message:     msg,
	}
}
