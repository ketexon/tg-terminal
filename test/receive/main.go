package main

import (
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

	fmt.Println("Reading messages...")
	for {
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
	}
}
