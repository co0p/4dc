package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/term"

	"pomodoro/internal/timer"
	"pomodoro/internal/ui"
)

func main() {
	os.Exit(run())
}

func run() int {
	fd := int(os.Stdin.Fd())

	// Only enter raw mode when stdin is an actual terminal.
	if term.IsTerminal(fd) {
		oldState, err := term.MakeRaw(fd)
		if err == nil {
			defer term.Restore(fd, oldState)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		cancel()
	}()

	ui.RunCycle(ctx, os.Stdin, os.Stdout, timer.NewCycle(
		25*time.Minute,
		5*time.Minute,
		15*time.Minute,
	))
	return 0
}
