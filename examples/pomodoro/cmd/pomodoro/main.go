package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"os/signal"
	"syscall"

	"github.com/co0p/4dc/examples/pomodoro/assets"
	"github.com/co0p/4dc/examples/pomodoro/internal/app"
	"github.com/co0p/4dc/examples/pomodoro/internal/tray"
)

var (
	flagVersion = flag.Bool("version", false, "print version and exit")
	flagSmoke   = flag.Bool("smoke", false, "run smoke startup and exit")
)

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Println("pomodoro-demo 0.1.0")
		return
	}

	// simple human-friendly logger
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("pomodoro: ")

	// use short durations for local demo default; domain durations are configurable
	a := app.New(25*time.Minute, 5*time.Minute)

	log.Println("starting application")

	if *flagSmoke {
		// initialize and immediately shutdown to validate startup
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := a.Shutdown(ctx); err != nil {
			log.Printf("smoke shutdown error: %v", err)
			os.Exit(2)
		}
		log.Println("smoke OK")
		return
	}

	// construct tray (no icon for now)
	t := tray.NewSystray(a, assets.Icon())

	// handle OS signals for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Println("signal received, shutting down")
		cancel()
	}()

	if err := t.Run(ctx); err != nil {
		log.Printf("tray.Run error: %v", err)
		os.Exit(1)
	}
	log.Println("exited")
}
