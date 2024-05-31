package main

import (
	"fmt"
	"time"
)

type DataServer struct {
	quitCh chan struct{}
	msgCh  chan string
}

func newServer() *DataServer {
	return &DataServer{
		quitCh: make(chan struct{}),
		msgCh:  make(chan string, 128),
	}
}

func (s *DataServer) start() {
	fmt.Println("Server is running")
}

func (s *DataServer) sendMessage(msg string) {
	s.msgCh <- msg
}

func (s *DataServer) quit() {
	s.quitCh <- struct{}{}
}

func (s *DataServer) loop() {
mainLoop:
	for {
		select {
		case <-s.quitCh:
			break mainLoop
		case msg := <-s.msgCh:
			s.handleMessage(msg)
		default:
		}
	}
	fmt.Println("server is shuting down gracefully")
}

func (s *DataServer) handleMessage(msg string) {
	fmt.Println("we received a message:", msg)
}

func syncy() {
	serve := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		serve.quit()
	}()

	serve.start()
}
