package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"ketexon/tg/proto"
	"net"
	"os"
)

func main() {
	addr := "localhost:8080"
	fmt.Printf("Connecting to server at %s...\n", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to server. Type a message and press Enter to send. Ctrl+C to quit.")
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		var msg string
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				panic(scanner.Err())
			}
			break
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

		msg = scanner.Text()
		if msg == "" {
			continue
		}

		var test proto.Test
		err := json.Unmarshal([]byte(msg), &test)
		if err != nil {
			fmt.Printf("Could not unmarshal message: %s\n", err)
			continue
		}

		fmt.Fprint(conn, msg)
	}
	fmt.Println("Shutting down...")
}
