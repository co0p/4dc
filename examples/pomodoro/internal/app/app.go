package app

import (
	"context"
	"sync"
	"time"
)

type State string

const (
	StateIdle            State = "Idle"
	StatePomodoroRunning State = "PomodoroRunning"
	StateBreakRunning    State = "BreakRunning"
)

type App interface {
	StartPomodoro()
	StartBreak()
	Shutdown(ctx context.Context) error
	OnStateChange(fn func(State))
	State() State
}

type timerApp struct {
	mu               sync.Mutex
	state            State
	cancelTimer      context.CancelFunc
	onStateChange    func(State)
	pomodoroDuration time.Duration
	breakDuration    time.Duration
}

// New creates a new App instance with configurable durations (useful for tests).
func New(pomodoro, brk time.Duration) App {
	return &timerApp{
		state:            StateIdle,
		pomodoroDuration: pomodoro,
		breakDuration:    brk,
	}
}

func (a *timerApp) setState(s State) {
	a.mu.Lock()
	a.state = s
	cb := a.onStateChange
	a.mu.Unlock()
	if cb != nil {
		cb(s)
	}
}

func (a *timerApp) State() State {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.state
}

func (a *timerApp) OnStateChange(fn func(State)) {
	a.mu.Lock()
	a.onStateChange = fn
	a.mu.Unlock()
}

func (a *timerApp) cancelExistingTimer() {
	if a.cancelTimer != nil {
		a.cancelTimer()
		a.cancelTimer = nil
	}
}

func (a *timerApp) StartPomodoro() {
	a.mu.Lock()
	if a.state == StatePomodoroRunning {
		a.mu.Unlock()
		return
	}
	a.cancelExistingTimer()
	ctx, cancel := context.WithCancel(context.Background())
	a.cancelTimer = cancel
	a.state = StatePomodoroRunning
	cb := a.onStateChange
	a.mu.Unlock()

	if cb != nil {
		cb(StatePomodoroRunning)
	}

	go func() {
		select {
		case <-time.After(a.pomodoroDuration):
			a.mu.Lock()
			a.cancelTimer = nil
			a.state = StateIdle
			cb := a.onStateChange
			a.mu.Unlock()
			if cb != nil {
				cb(StateIdle)
			}
		case <-ctx.Done():
			return
		}
	}()
}

func (a *timerApp) StartBreak() {
	a.mu.Lock()
	if a.state == StateBreakRunning {
		a.mu.Unlock()
		return
	}
	a.cancelExistingTimer()
	ctx, cancel := context.WithCancel(context.Background())
	a.cancelTimer = cancel
	a.state = StateBreakRunning
	cb := a.onStateChange
	a.mu.Unlock()

	if cb != nil {
		cb(StateBreakRunning)
	}

	go func() {
		select {
		case <-time.After(a.breakDuration):
			a.mu.Lock()
			a.cancelTimer = nil
			a.state = StateIdle
			cb := a.onStateChange
			a.mu.Unlock()
			if cb != nil {
				cb(StateIdle)
			}
		case <-ctx.Done():
			return
		}
	}()
}

func (a *timerApp) Shutdown(ctx context.Context) error {
	a.mu.Lock()
	a.cancelExistingTimer()
	a.state = StateIdle
	cb := a.onStateChange
	a.mu.Unlock()
	if cb != nil {
		cb(StateIdle)
	}
	return nil
}
