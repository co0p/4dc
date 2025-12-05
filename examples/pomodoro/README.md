# Pomodoro demo

This is a tiny demo showing a minimal app + tray wiring for the `4dc` examples. It demonstrates:

- A small domain `internal/app` with a simple timer state machine and tests.
- A platform `internal/tray` package using a systray implementation and a mock for tests.
- A CLI `cmd/pomodoro` with `--smoke` and graceful shutdown on `Quit`.

Quick build & run

```
cd examples/pomodoro
go build -o bin/pomodoro ./cmd/pomodoro
./bin/pomodoro
```

Smoke test (CI / local quick check):

```
./bin/pomodoro --smoke
```

Note about build artifacts
-------------------------

- **Do not commit build artifacts.** The `bin/` directory is used for local build outputs; avoid committing compiled binaries into the repository. To reproduce the demo locally, build the binary as shown above. For published release artifacts prefer placing them in a `releases/` directory or attaching them to a release.

Recommended quick build command:

```
cd examples/pomodoro
go build -o bin/pomodoro ./cmd/pomodoro
```

Acceptance (manual):

- The tray/menu shows `Pomodoro`, `Break`, and `Quit`.
- Clicking `Pomodoro` or `Break` triggers the app state change (check logs).
- Clicking `Quit` performs a graceful shutdown and exits the app.
- A minimal red-circle icon is shown in the tray.
 - While a Pomodoro or Break is running, the tray shows a concise remaining-time label in minutes (for example `25m` for a just-started Pomodoro). The label updates at a coarse cadence (approximately every 10s) and returns to the default tray state when the session finishes or is cancelled.

Replacing the icon

Currently the demo generates a simple 32×32 red-circle PNG at runtime (no external asset required). To use a custom icon instead, you have two options:

1. Provide a PNG file and update `assets/icon.go` to `//go:embed` it (recommended for a real icon):

	- Place your PNG at `examples/pomodoro/assets/icon.png`.
	- Update `assets/icon.go` to embed and return that file (replace the runtime generator).

2. Or provide a base64-encoded PNG in `examples/pomodoro/assets/icon.txt` and adjust `assets/icon.go` to decode it on startup.

Notes for macOS

- `systray` must be run on the main OS thread on macOS. This demo calls `systray.Run` from `main()` to satisfy that requirement.
- The demo was exercised on macOS; behavior on other platforms may vary.

Testing & smoke

- Unit tests: `go test ./...` (the example includes tests for `internal/app` and `internal/tray` mock).
- Smoke run: `./bin/pomodoro --smoke` initializes subsystems and exits immediately (used for CI/local checks).

Notes for contributors
----------------------

- We document a small set of code conventions and runtime constraints in `examples/pomodoro/docs/`.
- See `ADR-2025-12-05-receiver-naming-and-docs.md` for preferred receiver naming and godoc comment style (short receiver names, godoc sentences starting with the symbol name).
- There is also a short note about systray threading in `internal/tray/doc.go`; `systray.Run` must be called on the main OS thread on macOS. The `internal/tray` package wires the `TitleUpdater` but keep thread-safety in mind when moving calls that interact with the OS.

PR / branch

If you want me to open a focused PR with these example changes I will create branch `examples/pomodoro-demo` and push the `examples/pomodoro` changes.

Troubleshooting

- If the tray doesn't appear on macOS, ensure the binary is running and check logs printed to the terminal.

More

See `internal/app`, `internal/tray`, and `cmd/pomodoro` for the implementation.