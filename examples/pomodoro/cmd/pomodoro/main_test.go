package main_test

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"
)

// buildBinary compiles the pomodoro binary into a temp file and returns the path.
func buildBinary(t *testing.T) string {
	t.Helper()
	bin, err := os.CreateTemp(t.TempDir(), "pomodoro-test-*")
	if err != nil {
		t.Fatal(err)
	}
	bin.Close()
	build := exec.Command("go", "build", "-o", bin.Name(), "./cmd/pomodoro")
	build.Dir = "../.." // examples/pomodoro
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, out)
	}
	return bin.Name()
}

// accumulate reads from r in the background and sends chunks to the returned
// channel. The channel is closed when r returns an error (including EOF).
func accumulate(r io.Reader) <-chan string {
	ch := make(chan string, 256)
	go func() {
		defer close(ch)
		buf := make([]byte, 256)
		for {
			n, err := r.Read(buf)
			if n > 0 {
				ch <- string(buf[:n])
			}
			if err != nil {
				return
			}
		}
	}()
	return ch
}

// waitFor drains chunks from ch into *acc until acc contains substr or timeout.
func waitFor(ch <-chan string, acc *string, substr string, timeout time.Duration) bool {
	deadline := time.After(timeout)
	for {
		if strings.Contains(*acc, substr) {
			return true
		}
		select {
		case chunk, ok := <-ch:
			if !ok {
				return strings.Contains(*acc, substr)
			}
			*acc += chunk
		case <-deadline:
			return strings.Contains(*acc, substr)
		}
	}
}

// TestCtrlCExitsCleanly builds the binary, starts it, waits until the idle
// message appears on stdout, sends SIGINT, and asserts exit code 0.
func TestCtrlCExitsCleanly(t *testing.T) {
	cmd := exec.Command(buildBinary(t))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("start failed: %v", err)
	}

	ch := accumulate(stdout)
	var acc string
	if !waitFor(ch, &acc, "No session in progress", 3*time.Second) {
		cmd.Process.Kill()
		t.Fatalf("idle message did not appear within 3s, got: %q", acc)
	}

	if err := cmd.Process.Signal(syscall.SIGINT); err != nil {
		t.Fatalf("signal failed: %v", err)
	}

	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("process exited with error: %v", err)
		}
	case <-time.After(3 * time.Second):
		cmd.Process.Kill()
		t.Fatal("process did not exit within 3s after SIGINT")
	}
}

// TestCtrlCByteExitsCleanly sends raw byte 0x03 (Ctrl+C in raw mode) to stdin
// and asserts the process exits cleanly. This covers the case where MakeRaw has
// disabled ISIG so the terminal no longer generates SIGINT for Ctrl+C.
func TestCtrlCByteExitsCleanly(t *testing.T) {
	cmd := exec.Command(buildBinary(t))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("start failed: %v", err)
	}

	ch := accumulate(stdout)
	var acc string
	if !waitFor(ch, &acc, "No session in progress", 3*time.Second) {
		cmd.Process.Kill()
		t.Fatalf("idle message did not appear within 3s, got: %q", acc)
	}

	// Send the raw Ctrl+C byte that raw mode forwards instead of SIGINT.
	if _, err := stdin.Write([]byte{0x03}); err != nil {
		cmd.Process.Kill()
		t.Fatalf("failed to write 0x03 to stdin: %v", err)
	}

	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("process exited with error: %v", err)
		}
	case <-time.After(3 * time.Second):
		cmd.Process.Kill()
		t.Fatal("process did not exit within 3s after 0x03")
	}
}
func TestEnterStartsCountdown(t *testing.T) {
	cmd := exec.Command(buildBinary(t))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("start failed: %v", err)
	}

	ch := accumulate(stdout)
	var acc string

	if !waitFor(ch, &acc, "No session in progress", 3*time.Second) {
		cmd.Process.Kill()
		t.Fatalf("idle message did not appear within 3s, got: %q", acc)
	}

	if _, err := stdin.Write([]byte("\n")); err != nil {
		cmd.Process.Kill()
		t.Fatalf("failed to write to stdin: %v", err)
	}

	if !waitFor(ch, &acc, "25:00", 3*time.Second) {
		cmd.Process.Kill()
		t.Fatalf("countdown did not appear within 3s after Enter, got: %q", acc)
	}

	cmd.Process.Signal(syscall.SIGINT)
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		cmd.Process.Kill()
		t.Fatal("process did not exit after SIGINT")
	}
}
