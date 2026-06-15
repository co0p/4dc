package timer

import "time"

// IntervalKind identifies the type of a Pomodoro cycle interval.
type IntervalKind int

const (
	Pomodoro   IntervalKind = iota
	ShortBreak IntervalKind = iota
	LongBreak  IntervalKind = iota
)

// Interval is a single entry in a Pomodoro cycle.
// Number is 1–4 for Pomodoro intervals; 0 for breaks.
type Interval struct {
	Kind    IntervalKind
	Number  int
	Session Session
}

// NewCycle returns the 8-interval standard Pomodoro sequence:
// Pomodoro 1, Short break, Pomodoro 2, Short break, Pomodoro 3,
// Short break, Pomodoro 4, Long break.
// All Session fields use the production wall-clock ticker.
func NewCycle(work, short, long time.Duration) []Interval {
	return []Interval{
		{Kind: Pomodoro, Number: 1, Session: NewSession(work)},
		{Kind: ShortBreak, Number: 0, Session: NewSession(short)},
		{Kind: Pomodoro, Number: 2, Session: NewSession(work)},
		{Kind: ShortBreak, Number: 0, Session: NewSession(short)},
		{Kind: Pomodoro, Number: 3, Session: NewSession(work)},
		{Kind: ShortBreak, Number: 0, Session: NewSession(short)},
		{Kind: Pomodoro, Number: 4, Session: NewSession(work)},
		{Kind: LongBreak, Number: 0, Session: NewSession(long)},
	}
}
