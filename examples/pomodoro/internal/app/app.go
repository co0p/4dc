// Package app contains the core pomodoro timer domain logic. It exposes a
// small interface used by the demo UI to start sessions, observe state
// transitions, and query remaining time.
package app

import (
	"context"
	"sync"
	"time"
)

// State represents the observable lifecycle state of the pomodoro app.
//
// Typical values are `StateIdle`, `StatePomodoroRunning` and
// `StateBreakRunning`.
type State string

const (
	StateIdle            State = "Idle"
	StatePomodoroRunning State = "PomodoroRunning"
	StateBreakRunning    State = "BreakRunning"
)

// App is the minimal domain API the demo UI uses. It allows starting a
// pomodoro or break, shutting down the app, subscribing to state changes,
// and querying the current state and remaining time.
type App interface {
	StartPomodoro()
	StartBreak()
	Shutdown(ctx context.Context) error
	OnStateChange(fn func(State))
	// SubscribeStateChange registers a listener for state changes and returns an
	// unsubscribe function. The unsubscribe function is safe to call multiple
	// times and may be called from any goroutine.
	SubscribeStateChange(fn func(State)) func()
	State() State
	Remaining() time.Duration
}

type timerApp struct {
	mu          sync.Mutex
	state       State
	cancelTimer context.CancelFunc
	// subscribers holds active state-change listeners keyed by id.
	subscribers      map[int]func(State)
	nextSubID        int
	pomodoroDuration time.Duration
	breakDuration    time.Duration
	end              time.Time
}

// New creates a new App instance. Optionally pass two durations: pomodoro, break.
// Examples:
//
//	New() // uses defaults
//	New(10*time.Millisecond, 5*time.Millisecond) // test-friendly durations
func New(durations ...time.Duration) App {
	pom := 25 * time.Minute
	brk := 5 * time.Minute
	if len(durations) >= 2 {
		pom = durations[0]
		brk = durations[1]
	}
	return &timerApp{
		state:            StateIdle,
		pomodoroDuration: pom,
		breakDuration:    brk,
	}
}

// State returns the current lifecycle state of the app.
func (t *timerApp) State() State {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.state
}

// OnStateChange registers a callback to be invoked on state transitions.
// The callback is invoked asynchronously. For an unsubscribe function
// use `SubscribeStateChange` which returns a cleanup function.
func (t *timerApp) OnStateChange(fn func(State)) {
	// Keep compatibility: subscribe and discard the unsubscribe function.
	t.SubscribeStateChange(fn)
}

// SubscribeStateChange registers a listener for state changes and returns
// an unsubscribe function. The unsubscribe function is safe to call
// multiple times and may be called from any goroutine.
func (t *timerApp) SubscribeStateChange(fn func(State)) func() {
	t.mu.Lock()
	if t.subscribers == nil {
		t.subscribers = make(map[int]func(State))
	}
	id := t.nextSubID
	t.nextSubID++
	t.subscribers[id] = fn
	t.mu.Unlock()

	return func() {
		t.mu.Lock()
		delete(t.subscribers, id)
		t.mu.Unlock()
	}
}

func (t *timerApp) notifySubscribers(s State) {
	t.mu.Lock()
	// copy subscribers to avoid holding lock while calling callbacks
	subs := make([]func(State), 0, len(t.subscribers))
	for _, fn := range t.subscribers {
		subs = append(subs, fn)
	}
	t.mu.Unlock()

	for _, fn := range subs {
		if fn != nil {
			fn(s)
		}
	}
}

func (t *timerApp) cancelExistingTimer() {
	if t.cancelTimer != nil {
		t.cancelTimer()
		t.cancelTimer = nil
		t.end = time.Time{}
	}
}

// Remaining returns the remaining duration for the current running
// session (pomodoro or break). It returns zero when no session is
// active or when the remaining time has elapsed.
func (t *timerApp) Remaining() time.Duration {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.state != StatePomodoroRunning && t.state != StateBreakRunning {
		return 0
	}
	if t.end.IsZero() {
		return 0
	}
	d := time.Until(t.end)
	if d <= 0 {
		return 0
	}
	return d
}

// StartPomodoro begins a pomodoro session. If a pomodoro is already
// running this is a no-op.
func (t *timerApp) StartPomodoro() {
	t.mu.Lock()
	if t.state == StatePomodoroRunning {
		t.mu.Unlock()
		return
	}
	t.cancelExistingTimer()
	ctx, cancel := context.WithCancel(context.Background())
	t.cancelTimer = cancel
	t.end = time.Now().Add(t.pomodoroDuration)
	t.state = StatePomodoroRunning
	t.mu.Unlock()

	t.notifySubscribers(StatePomodoroRunning)

	go func() {
		select {
		case <-time.After(t.pomodoroDuration):
			t.mu.Lock()
			t.cancelTimer = nil
			t.state = StateIdle
			t.mu.Unlock()
			t.notifySubscribers(StateIdle)
		case <-ctx.Done():
			return
		}
	}()
}

// StartBreak begins a break session. If a break is already running this
// is a no-op.
func (t *timerApp) StartBreak() {
	t.mu.Lock()
	if t.state == StateBreakRunning {
		t.mu.Unlock()
		return
	}
	t.cancelExistingTimer()
	ctx, cancel := context.WithCancel(context.Background())
	t.cancelTimer = cancel
	t.end = time.Now().Add(t.breakDuration)
	t.state = StateBreakRunning
	t.mu.Unlock()

	t.notifySubscribers(StateBreakRunning)

	go func() {
		select {
		case <-time.After(t.breakDuration):
			t.mu.Lock()
			t.cancelTimer = nil
			t.state = StateIdle
			t.mu.Unlock()
			t.notifySubscribers(StateIdle)
		case <-ctx.Done():
			return
		}
	}()
}

// Shutdown stops any active session, transitions the app to idle, and
// performs any necessary cleanup. The provided context may be used to
// bound shutdown operations (currently unused by the simple demo
// implementation).
func (t *timerApp) Shutdown(ctx context.Context) error {
	t.mu.Lock()
	t.cancelExistingTimer()
	t.state = StateIdle
	t.mu.Unlock()
	t.notifySubscribers(StateIdle)
	return nil
}
