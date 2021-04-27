package pong

import "time"

// Pong is payload sent to client
type Pong struct {
	CurrentTime time.Time
	Message     string
}

// NewPong creates a new Pong
func NewPong(msg string) Pong {
	return Pong{
		CurrentTime: time.Now().UTC(),
		Message:     msg,
	}
}
