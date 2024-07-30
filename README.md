multislog
=========

A package that wraps, in a single slog.Handler, an array of other handlers.

Example usage:
```go
package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/alchemy/multislog"
)

func main() {
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	stdoutHandler := slog.NewTextHandler(os.Stdout, opts)
	sb := &strings.Builder{}
	stringHandler := slog.NewTextHandler(sb, opts)
	handler := multislog.NewHandler(stdoutHandler, stringHandler)
	logger := slog.New(handler)
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
}
```