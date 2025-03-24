package main

import (
	"context"
	"ketexon/tg/message"
	pb "ketexon/tg/proto"
	"ketexon/tg/tui"
	vshell "ketexon/tg/tui/vshell"
	"ketexon/tg/vos"
	"log/slog"
	"os"
	"os/signal"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	logPath := "log.txt"
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	logger := slog.New(
		slog.NewTextHandler(
			logFile,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			},
		),
	)
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())

	// create a new message server
	server := message.NewServer()

	// the TUI
	var p *tea.Program = nil

	// subscribe
	server.Subscribe(func(msg *pb.GameMessage) {
		switch msg.Content.(type) {
		case *pb.GameMessage_Ping:
			server.SendMessage(&pb.TerminalMessage{
				Content: &pb.TerminalMessage_Ping{},
			})
		case *pb.GameMessage_Login:
			{
				slog.Info("Received login message", "computerId", msg.GetLogin().GetComputerId())
				// try to get VOS from the message
				if vos, ok := vos.VOSMap[msg.GetLogin().GetComputerId()]; ok {
					p.Send(tui.SwitchModel{
						NextModel: vshell.New(vos),
					})
				} else {
					// send error message
					slog.Error("VOS not found for computerId", "computerId", msg.GetLogin().GetComputerId())
					server.SendMessage(&pb.TerminalMessage{
						Content: &pb.TerminalMessage_Error{
							Error: &pb.Error{
								Message: "VOS not found for computerId",
							},
						},
					})
				}
			}
		}
	})

	if err := server.Start(ctx); err != nil {
		panic(err)
	}

	// create a new TUI
	var startModel tea.Model = nil
	if vos, ok := vos.VOSMap["user1"]; ok {
		startModel = vshell.New(vos)
	} else {
		slog.Error("Could not find default VOS")
	}

	p = tea.NewProgram(
		tui.NewTUI(
			server,
			startModel,
		),
		tea.WithAltScreen(),
	)
	go func() {
		if _, err := p.Run(); err != nil {
			slog.Error("Error starting TUI", "err", err)
		}
		cancel()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		sig := <-sigCh
		if sig != nil {
			slog.Info("Received interrupt signal.", "sig", sig)
			cancel()
		}
	}()

	<-ctx.Done()
	close(sigCh)
	if p != nil {
		p.Quit()
	}

	slog.Info("Shutting down...")
	os.Exit(0)
}
