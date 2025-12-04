package tray

import (
	"context"
	"log"
	"time"

	"github.com/co0p/4dc/examples/pomodoro/internal/app"
	"github.com/getlantern/systray"
)

type systrayImpl struct {
	app  app.App
	icon []byte
}

func newSystrayImpl(a app.App, icon []byte) Tray {
	return &systrayImpl{app: a, icon: icon}
}

func (s *systrayImpl) Run(ctx context.Context) error {
	done := make(chan struct{})

	// start systray in a goroutine because Run blocks
	go func() {
		systray.Run(func() {
			if len(s.icon) > 0 {
				systray.SetIcon(s.icon)
			}
			mPom := systray.AddMenuItem("Pomodoro", "Start Pomodoro")
			mBreak := systray.AddMenuItem("Break", "Start Break")
			systray.AddSeparator()
			mQuit := systray.AddMenuItem("Quit", "Quit the app")

			// listen for menu clicks
			go func() {
				for range mPom.ClickedCh {
					log.Println("action=StartPomodoro")
					s.app.StartPomodoro()
				}
			}()
			go func() {
				for range mBreak.ClickedCh {
					log.Println("action=StartBreak")
					s.app.StartBreak()
				}
			}()
			go func() {
				for range mQuit.ClickedCh {
					log.Println("action=Quit")
					// call shutdown synchronously with a timeout
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					_ = s.app.Shutdown(ctx)
					cancel()
					systray.Quit()
				}
			}()
		}, func() {
			close(done)
		})
	}()

	select {
	case <-ctx.Done():
		systray.Quit()
		<-done
		return ctx.Err()
	case <-done:
		return nil
	}
}

func (s *systrayImpl) Close() error {
	systray.Quit()
	return nil
}
