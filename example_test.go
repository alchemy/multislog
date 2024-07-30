package multislog_test

import (
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/alchemy/multislog"
)

func Example() {
	t, err := time.Parse("2006-01-02 15:04:05", "2009-11-10 23:00:00")
	if err != nil {
		log.Fatalln(err)
	}
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Time(slog.TimeKey, t)
			}
			return a
		},
	}

	stdoutHandler := slog.NewTextHandler(os.Stdout, opts)
	sb := &strings.Builder{}
	stringHandler := slog.NewTextHandler(sb, opts)
	handler := multislog.NewHandler(stdoutHandler, stringHandler)
	logger := slog.New(handler)
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	// Output:
	// time=2009-11-10T23:00:00.000Z level=INFO msg="info message"
	// time=2009-11-10T23:00:00.000Z level=WARN msg="warn message"
	// time=2009-11-10T23:00:00.000Z level=ERROR msg="error message"
}
