package main

import (
	"encoding/binary"
	"fmt"
	pb "ketexon/tg/proto"
	"net"

	"google.golang.org/protobuf/proto"
)

func main() {
	addr := "localhost:8080"
	fmt.Printf("Connecting to server at %s...\n", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	msg := &pb.GameMessage{
		Content: &pb.GameMessage_Login{
			Login: &pb.Login{
				ComputerId: "user1",
			},
		},
	}
	out, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	out = append(make([]byte, 4), out...) // Append a null byte to indicate end of message
	binary.BigEndian.PutUint32(out[:4], uint32(len(out)))
	conn.Write(out)

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	res := &pb.TerminalMessage{}
	err = proto.Unmarshal(buf[4:n], res)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received response: %s\n", res.String())
	fmt.Println("Shutting down...")
}
