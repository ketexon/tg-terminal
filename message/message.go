package message

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	pb "ketexon/tg/proto"
	"log/slog"
	"net"

	"google.golang.org/protobuf/proto"
)

const Port = 8080

type GameMessageCallback func(msg *pb.GameMessage)

type Server struct {
	ln     net.Listener
	ch     chan *pb.TerminalMessage
	subs   []GameMessageCallback
	logger *slog.Logger
}

func NewServer() *Server {
	return &Server{
		ln:     nil,
		ch:     make(chan *pb.TerminalMessage, 10),
		subs:   make([]GameMessageCallback, 0),
		logger: slog.Default().WithGroup("message"),
	}
}

func (s *Server) Start(ctx context.Context) error {
	addr := fmt.Sprintf("localhost:%d", Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	slog.Info("Server started", "addr", addr)
	s.ln = ln
	context.AfterFunc(ctx, func() {
		s.logger.Info("Shutting down server...")
		s.ln.Close()
	})
	go s.connectRoutine()
	return nil
}

func (s *Server) connectRoutine() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				s.logger.Info("Server closed, stopping accept loop")
				break
			}
			s.logger.Error("Could not connect", "err", err)
			continue
		}
		s.logger.Info("Connected to client", "conn", conn.RemoteAddr())
		s.messageRoutine(conn)
	}
}

const messageBuffer = 16384 // 16KB

func (s *Server) messageRoutine(conn net.Conn) {
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())

	go s.sendMessageRoutine(conn, ctx)

	buf := make([]byte, messageBuffer)
messageLoop:
	for {
		var err error
		// read new message
		bytes_read := 0
		for bytes_read < 4 {
			n, err := conn.Read(buf[bytes_read:4])
			if err == io.EOF {
				break messageLoop
			}
			if err != nil {
				s.logger.Error("Errror reading connection", "conn", conn.RemoteAddr(), "err", err)
				break messageLoop
			}
			bytes_read += n
		}

		msg_length := int(binary.BigEndian.Uint32(buf[:4]))
		if msg_length > messageBuffer {
			s.logger.Error("Message too large", "length", msg_length)
			break messageLoop
		}

		for bytes_read < msg_length {
			n, err := conn.Read(buf[bytes_read:])
			if err == io.EOF {
				break
			}
			if err != nil {
				s.logger.Error("Errror reading connection", "conn", conn.RemoteAddr(), "err", err)
				break messageLoop
			}
			bytes_read += n
		}

		msg := &pb.GameMessage{}
		err = proto.Unmarshal(buf[4:msg_length], msg)
		if err != nil {
			s.logger.Error("Could not unmarshal message", "conn", conn.RemoteAddr(), "err", err)
			continue
		}
		s.logger.Info("Received message", "conn", conn.RemoteAddr(), "msg", msg.String())
		s.publishMessage(msg)
	}
	s.logger.Info("Client disconnected", "conn", conn.RemoteAddr())
	cancel()
}

func (s *Server) Subscribe(callback GameMessageCallback) {
	s.subs = append(s.subs, callback)
}

func (s *Server) publishMessage(msg *pb.GameMessage) {
	for _, cb := range s.subs {
		cb(msg)
	}
}

func (s *Server) sendMessageRoutine(conn net.Conn, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-s.ch:
			out, err := proto.Marshal(msg)
			if err != nil {
				s.logger.Error("Could not marshal message", "err", err)
				continue
			}
			out = append(make([]byte, 4), out...)
			binary.BigEndian.PutUint32(out[:4], uint32(len(out)))
			conn.Write(out)
		}
	}
}

func (s *Server) SendMessage(msg *pb.TerminalMessage) {
	s.ch <- msg
}
