# Pomodoro — terminal PRD

A minimal command-line Pomodoro timer. No GUI, no system tray, no external services. Runs anywhere a shell runs.

---

## Build

Requires Go 1.25+.

```bash
go build -o pomodoro ./cmd/pomodoro
./pomodoro
```

## Usage

Press **Enter** or **Space** to start the Pomodoro cycle.
The display updates every second, showing the current interval:

```
[Pomodoro 1/4] 24:59 remaining
[Short break]  04:59 remaining
[Pomodoro 2/4] 24:59 remaining
[Short break]  04:59 remaining
[Pomodoro 3/4] 24:59 remaining
[Short break]  04:59 remaining
[Pomodoro 4/4] 24:59 remaining
[Long break]   14:59 remaining
```

A bell sounds at the end of each interval. All 8 intervals advance automatically —
no key press needed between them. After the long break the app returns to the idle screen.

Press **Ctrl-C** at any time to quit.

---

## Problem

Context-switching is expensive. The Pomodoro Technique enforces focused work intervals and mandatory breaks. Existing tools are heavyweight or require a display server. A developer working over SSH, in a container, or inside a tmux session needs a timer that lives in the terminal.

---

## Goal

A single binary that runs a standard Pomodoro cycle in the terminal. Prefer standard-library primitives; third-party libraries are acceptable if they are small, well-maintained, and add clear value (e.g. raw-mode terminal input). Avoid large frameworks or transitive dependency trees.

---

## Technique

Standard Pomodoro Technique — no customisation required for v1:

| Interval | Duration |
|----------|----------|
| Work (Pomodoro) | 25 min |
| Short break | 5 min |
| Long break (every 4th) | 15 min |

Cycle: work → short break → work → short break → work → short break → work → long break → repeat.

---

## User interface

All interaction happens in the terminal.

**Display** — a single line updated in place:

```
[Pomodoro 3/4]  18:32 remaining
```

**Controls** — keyboard only:

| Key | Action |
|-----|--------|
| `Enter` or `Space` | Skip to next interval |
| `q` or `Ctrl-C` | Quit |

**Session end** — print a summary line and exit:

```
Session complete. 4 pomodoros, 100 min focused.
```

**Notification** — print a visible bell (`\a`) and a one-line message when an interval ends. No OS notification APIs.

---

## Acceptance criteria

- [ ] Starts with no arguments: `pomodoro`
- [ ] Counts down correctly; display updates every second
- [ ] Advances automatically to the next interval when the timer reaches zero
- [ ] Long break fires after every 4th completed Pomodoro
- [ ] Skip and quit keys work reliably in a standard terminal
- [ ] Exits cleanly on `Ctrl-C`; restores terminal state
- [ ] Ships as a single self-contained binary (no runtime installation required)
- [ ] Third-party build dependencies, if any, are minimal and vendored or pinned
- [ ] No GUI, no network, no config file required

---

## Out of scope (v1)

- Custom durations
- Persistence / history
- OS notifications
- Sound files
- Configuration file
- Windows support
