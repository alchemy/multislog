package multislog

import (
	"log/slog"
	"strings"
	"testing"
)

func TestHandler1(t *testing.T) {
	sb1 := &strings.Builder{}
	sb2 := &strings.Builder{}
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	h1 := slog.NewTextHandler(sb1, opts)
	h2 := slog.NewTextHandler(sb2, opts)
	h := NewHandler(h1, h2)
	logger := slog.New(h)
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	if sb1.String() != sb2.String() {
		t.Fatalf("handlers produced different output:\n  h1: %s\n  h2: %s", sb1, sb2)
	}
	if n := strings.Count(sb1.String(), "\n"); n != 3 {
		t.Fatalf("expected 3 log lines, got %d: %s", n, sb1)
	}
}
