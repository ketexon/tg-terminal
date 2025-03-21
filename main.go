package main

import (
	"context"
	"fmt"
	"ketexon/tg/message"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	server := message.NewServer()
	err := server.Start(ctx)
	if err != nil {
		panic(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		sig := <-sigCh
		fmt.Printf("Received signal: %s. Quitting...\n", sig)
		cancel()
	}()

	<-ctx.Done()
	close(sigCh)

	fmt.Println("Shutting down...")
	os.Exit(0)
}
