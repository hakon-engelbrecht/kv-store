// Package server contains the tcp server, that can be used to
// serve at a specified port.
package server

import (
	"fmt"
	"log"
	"net"
)

// TCPServer serves a configurable port and configurable handler function.
type TCPServer struct {
	port       int
	handleFunc func(net.Conn)
}

// NewTCPServer creates a new instance of a tcp server and configures
// the port and handler function.
func NewTCPServer(port int, handleFunc func(net.Conn)) *TCPServer {
	return &TCPServer{
		port:       port,
		handleFunc: handleFunc,
	}
}

// Listen makes the server listen on its configred port.
// The server handles incoming connections with the handler
// function specified on construction.
func (s *TCPServer) Listen() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		return err
	}
	log.Printf("server listening on port %v...", s.port)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("could not accept connection: ", err)
		}
		log.Println("accepted connection")
		go s.handleFunc(conn)
	}
}
