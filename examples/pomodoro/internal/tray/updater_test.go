package tray

import (
	"context"
	"testing"
	"time"

	"github.com/co0p/4dc/examples/pomodoro/internal/app"
)

type fakeApp struct {
	rem   time.Duration
	cb    func(app.State)
	wired chan struct{}
}

func (f *fakeApp) StartPomodoro()                     {}
func (f *fakeApp) StartBreak()                        {}
func (f *fakeApp) Shutdown(ctx context.Context) error { return nil }
func (f *fakeApp) OnStateChange(fn func(app.State))   { f.cb = fn }
func (f *fakeApp) SubscribeStateChange(fn func(app.State)) func() {
	f.cb = fn
	// signal that subscription is wired
	if f.wired != nil {
		close(f.wired)
		f.wired = nil
	}
	return func() { f.cb = nil }
}
func (f *fakeApp) State() app.State         { return app.StateIdle }
func (f *fakeApp) Remaining() time.Duration { return f.rem }

// TestTitleUpdaterDeterministic verifies TitleUpdater updates the title
// immediately on transition to running, on ticks, and clears on idle.
func TestTitleUpdaterDeterministic(t *testing.T) {
	f := &fakeApp{rem: 5 * time.Minute, wired: make(chan struct{})}

	titleCh := make(chan string, 10)
	setTitle := func(s string) { titleCh <- s }
	clearTitle := func() { titleCh <- "CLEAR" }

	var currentTickCh chan time.Time
	newTicker := func(d time.Duration) (<-chan time.Time, func()) {
		ch := make(chan time.Time)
		currentTickCh = ch
		return ch, func() {}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u := NewTitleUpdater(f, setTitle, clearTitle, newTicker)
	go u.Run(ctx)

	// wait for subscription to be wired
	select {
	case <-f.wired:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("subscription not wired")
	}

	// transition to running -> immediate update
	f.rem = 5 * time.Minute
	f.cb(app.StatePomodoroRunning)

	select {
	case got := <-titleCh:
		if got != "5m" {
			t.Fatalf("expected initial title 5m, got %q", got)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout waiting for initial title")
	}

	// simulate a tick with updated remaining
	f.rem = 4 * time.Minute
	if currentTickCh == nil {
		t.Fatal("no tick channel available")
	}
	currentTickCh <- time.Now()
	select {
	case got := <-titleCh:
		if got != "4m" {
			t.Fatalf("expected tick title 4m, got %q", got)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("timeout waiting for tick title")
	}

	// transition to idle -> clear
	f.cb(app.StateIdle)
	select {
	case got := <-titleCh:
		if got != "CLEAR" {
			t.Fatalf("expected clear, got %q", got)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout waiting for clear title")
	}

	// stop updater
	cancel()
}

func TestTitleUpdaterStopUnsubscribes(t *testing.T) {
	f := &fakeApp{rem: 3 * time.Minute, wired: make(chan struct{})}

	titleCh := make(chan string, 10)
	setTitle := func(s string) { titleCh <- s }
	clearTitle := func() { titleCh <- "CLEAR" }

	newTicker := func(d time.Duration) (<-chan time.Time, func()) {
		ch := make(chan time.Time)
		return ch, func() {}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u := NewTitleUpdater(f, setTitle, clearTitle, newTicker)
	go u.Run(ctx)

	// wait for subscription to be wired
	select {
	case <-f.wired:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("subscription not wired")
	}

	// transition to running -> immediate update
	f.cb(app.StatePomodoroRunning)
	select {
	case got := <-titleCh:
		if got != "3m" {
			t.Fatalf("expected initial title 3m, got %q", got)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout waiting for initial title")
	}

	// stop updater -> should clear title and unsubscribe
	u.Stop()
	select {
	case got := <-titleCh:
		if got != "CLEAR" {
			t.Fatalf("expected clear after stop, got %q", got)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout waiting for clear after stop")
	}

	// attempt to trigger state change after stop; should not produce titles
	if f.cb != nil {
		f.cb(app.StatePomodoroRunning)
	}
	select {
	case got := <-titleCh:
		t.Fatalf("unexpected title after stop: %q", got)
	case <-time.After(100 * time.Millisecond):
		// no titles as expected
	}
}

func TestTitleUpdaterStartStopCycles(t *testing.T) {
	f := &fakeApp{rem: 2 * time.Minute, wired: make(chan struct{})}

	titleCh := make(chan string, 20)
	setTitle := func(s string) { titleCh <- s }
	clearTitle := func() { titleCh <- "CLEAR" }

	var currentTickCh chan time.Time
	newTicker := func(d time.Duration) (<-chan time.Time, func()) {
		ch := make(chan time.Time)
		currentTickCh = ch
		return ch, func() {}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u := NewTitleUpdater(f, setTitle, clearTitle, newTicker)
	go u.Run(ctx)

	// wait for subscription
	select {
	case <-f.wired:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("subscription not wired")
	}

	cycles := 3
	for i := 0; i < cycles; i++ {
		// start running -> immediate update
		f.rem = time.Duration(2-i) * time.Minute
		f.cb(app.StatePomodoroRunning)
		select {
		case got := <-titleCh:
			if got == "CLEAR" {
				t.Fatalf("unexpected CLEAR on start cycle %d", i)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timeout waiting for title on start cycle %d", i)
		}

		// simulate a tick
		if currentTickCh == nil {
			t.Fatalf("no current tick channel on cycle %d", i)
		}
		currentTickCh <- time.Now()
		select {
		case <-titleCh:
		case <-time.After(200 * time.Millisecond):
			t.Fatalf("timeout waiting for tick on cycle %d", i)
		}

		// stop -> should clear
		u.Stop()
		select {
		case got := <-titleCh:
			if got != "CLEAR" {
				t.Fatalf("expected CLEAR after stop on cycle %d, got %q", i, got)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timeout waiting for clear on cycle %d", i)
		}

		// ensure no title after external state change
		if f.cb != nil {
			f.cb(app.StatePomodoroRunning)
		}
		select {
		case got := <-titleCh:
			t.Fatalf("unexpected title after stop on cycle %d: %q", i, got)
		case <-time.After(100 * time.Millisecond):
			// good
		}

		// recreate updater for next cycle
		f.wired = make(chan struct{})
		u = NewTitleUpdater(f, setTitle, clearTitle, newTicker)
		go u.Run(ctx)

		// wait for reconnection
		select {
		case <-f.wired:
		case <-time.After(200 * time.Millisecond):
			t.Fatal("subscription not wired after recreating updater")
		}
	}

	cancel()
}
