package tray

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/co0p/4dc/examples/pomodoro/internal/app"
)

// titleUpdateInterval is the cadence at which the tray title is refreshed
// when a session is running.
const titleUpdateInterval = 10 * time.Second

// TitleUpdater manages the tray title lifecycle: it subscribes to app state
// changes, updates the title immediately on transitions to running, and
// periodically on a ticker. Stop() detaches subscriptions and stops the
// ticker. The updater accepts an injected ticker factory to make tests
// deterministic.
type TitleUpdater struct {
	app           app.App
	setTitle      func(string)
	clearTitle    func()
	tickerFactory func(d time.Duration) (<-chan time.Time, func())

	mu          sync.Mutex
	unsubscribe func()
	tickCh      <-chan time.Time
	stopTicker  func()
	running     bool
}

// NewTitleUpdater constructs a TitleUpdater. The tickerFactory returns a
// tick channel and a stopper function for the ticker.
func NewTitleUpdater(a app.App, setTitle func(string), clearTitle func(),
	tickerFactory func(d time.Duration) (<-chan time.Time, func())) *TitleUpdater {
	return &TitleUpdater{
		app:           a,
		setTitle:      setTitle,
		clearTitle:    clearTitle,
		tickerFactory: tickerFactory,
	}
}

// Run starts the updater loop and blocks until ctx is done. It subscribes to
// app state changes and ensures cleanup on exit.
func (t *TitleUpdater) Run(ctx context.Context) {
	stateCh := make(chan app.State, 1)

	unsub := t.app.SubscribeStateChange(func(s app.State) {
		select {
		case stateCh <- s:
		default:
		}
	})

	t.mu.Lock()
	t.unsubscribe = unsub
	t.mu.Unlock()

	for {
		select {
		case <-ctx.Done():
			t.mu.Lock()
			if t.running {
				if t.stopTicker != nil {
					t.stopTicker()
				}
				t.clearTitle()
			}
			if t.unsubscribe != nil {
				t.unsubscribe()
				t.unsubscribe = nil
			}
			t.mu.Unlock()
			return
		case s := <-stateCh:
			if s == app.StatePomodoroRunning || s == app.StateBreakRunning {
				// restart ticker on any transition to running
				t.mu.Lock()
				if t.stopTicker != nil {
					t.stopTicker()
				}
				ch, stopper := t.tickerFactory(titleUpdateInterval)
				t.tickCh = ch
				t.stopTicker = stopper
				t.running = true
				t.mu.Unlock()

				// immediate update
				rem := t.app.Remaining()
				mins := int(rem.Minutes())
				t.setTitle(formatMinutes(mins))
			} else if s == app.StateIdle {
				t.mu.Lock()
				if t.running {
					t.running = false
					if t.stopTicker != nil {
						t.stopTicker()
						t.stopTicker = nil
					}
					t.mu.Unlock()
					t.clearTitle()
				} else {
					t.mu.Unlock()
				}
			}
		case <-t.tickCh:
			t.mu.Lock()
			running := t.running
			t.mu.Unlock()
			if running {
				rem := t.app.Remaining()
				mins := int(rem.Minutes())
				t.setTitle(formatMinutes(mins))
			}
		}
	}
}

// Stop detaches the subscription and stops any running ticker. It is safe to
// call multiple times.
func (t *TitleUpdater) Stop() {
	t.mu.Lock()
	if t.unsubscribe != nil {
		t.unsubscribe()
		t.unsubscribe = nil
	}
	if t.stopTicker != nil {
		t.stopTicker()
		t.stopTicker = nil
	}
	running := t.running
	t.running = false
	t.mu.Unlock()

	if running {
		t.clearTitle()
	}
}

// ManageTitleUpdates kept for compatibility: it constructs a TitleUpdater
// and runs it until the context is cancelled.
func ManageTitleUpdates(ctx context.Context, a app.App, setTitle func(string), clearTitle func(),
	newTicker func(d time.Duration) (<-chan time.Time, func())) {
	u := NewTitleUpdater(a, setTitle, clearTitle, newTicker)
	u.Run(ctx)
}

func formatMinutes(m int) string {
	return fmt.Sprintf("%dm", m)
}
