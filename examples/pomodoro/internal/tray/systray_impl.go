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
	var updaterCancel context.CancelFunc

	// If the provided context is canceled, request systray to quit.
	go func() {
		<-ctx.Done()
		systray.Quit()
	}()

	// systray.Run must be called on the main OS thread on macOS. Call it
	// directly in this goroutine so callers that invoke Run from main()
	// satisfy that requirement.
	var u *TitleUpdater
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
				c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				_ = s.app.Shutdown(c)
				cancel()
				systray.Quit()
			}
		}()

		// Start title updater. It subscribes to app state changes and
		// periodically queries Remaining() to update the tray title.
		//
		// Note: `systray.Run` runs the systray event loop on the OS thread.
		// We construct the updater here and run it in a goroutine; the
		// provided `setTitle`/`clearTitle` call `systray.SetTitle` which is
		// implemented by the library and is safe to call from other
		// goroutines in this usage. If the underlying systray backend
		// requires main-thread calls for SetTitle, move `u.Run` onto the
		// systray callback thread instead.
		updaterCtx, cancel := context.WithCancel(context.Background())
		updaterCancel = cancel
		setTitle := func(t string) { systray.SetTitle(t) }
		clearTitle := func() { systray.SetTitle("") }
		newTicker := func(d time.Duration) (<-chan time.Time, func()) {
			t := time.NewTicker(d)
			return t.C, func() { t.Stop() }
		}
		u = NewTitleUpdater(s.app, setTitle, clearTitle, newTicker)
		go u.Run(updaterCtx)
	}, func() {
		if updaterCancel != nil {
			updaterCancel()
		}
		// If we constructed a TitleUpdater above, ensure it is stopped so
		// it detaches subscriptions and clears the title.
		// Note: the updater was run in a goroutine; Stop() is safe to call
		// concurrently.
		if u != nil {
			u.Stop()
		}
		close(done)
	})

	<-done
	return ctx.Err()
}

func (s *systrayImpl) Close() error {
	systray.Quit()
	return nil
}
