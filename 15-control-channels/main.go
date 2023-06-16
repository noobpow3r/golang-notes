package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // 0 bytes
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop() // block
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
			// do some stuff when we need to quit
		case msg := <-s.msgch:
			// do some stuff when we have a message
			s.handleMessage(msg)
		default:
		}
	}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message:", msg)
}

func main() {
	server := newServer()
	go server.start()

	// server.sendMessage("hey! do this ...")
	for i := 0; i < 100; i++ {
		server.sendMessage(fmt.Sprintf("handle this number %d", i))
	}

	time.Sleep(time.Second * 5)

}
