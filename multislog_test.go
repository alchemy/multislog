package multislog

import (
	"log/slog"
	"strings"
	"testing"

	"log"
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
		log.Fatalln(sb1, "differs from", sb2)
	}
	if len(strings.Split(sb1.String(), "\n")) != 4 {
		log.Fatalln(sb1, "must contain exactly 4 newline characters")
	}
}
