package main

import (
	"log"

	"pomodoro/internal/sampleactions"
	"pomodoro/pkg/tray"
)

func main() {
	log.Println("[app] starting Pomodoro tray app")
	adapter := tray.SystrayAdapter{}
	items := []tray.Item{
		{Title: "Start Sample", OnClick: sampleactions.StartSample},
		{Title: "Pause Sample", OnClick: sampleactions.PauseSample},
	}
	adapter.Show("Pomodoro", items)
	log.Println("[app] tray app exited")
}
