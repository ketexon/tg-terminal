package message

import (
	"context"
	"fmt"
	"io"
	"net"
)

const PORT = 8080

type Server struct {
	ln net.Listener
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(ctx context.Context) error {
	addr := fmt.Sprintf("localhost:%d", PORT)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	fmt.Printf("Server started on %s\n", addr)
	s.ln = ln
	context.AfterFunc(ctx, func() {
		fmt.Println("Shutting down server...")
		s.ln.Close()
	})
	go s.connectRoutine()
	return nil
}

func (s *Server) connectRoutine() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Connected to client %s\n", conn.RemoteAddr())
		go s.messageRoutine(conn)
	}
}

func (s *Server) messageRoutine(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if n == 0 {
			break
		}
		msg := string(buf[:n])
		fmt.Printf("Received message from %s: %s\n", conn.RemoteAddr(), msg)
	}
	fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
}
