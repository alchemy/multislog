package multislog

import (
	"context"
	"errors"
	"log/slog"
)

type handler []slog.Handler

// NewHandler creates a new [slog.Handler] which combines all the
// handlers passed as arguments in a single one.
func NewHandler(handlers ...slog.Handler) slog.Handler {
	return handler(handlers)
}

// Enabled implements the method of the slog.Handler interface.
func (h handler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

// Handle implements the method of the slog.Handler interface.
// It simply calls each underlying handler Handle method.
func (h handler) Handle(ctx context.Context, r slog.Record) error {
	var errs []error
	for i := range h {
		if !h[i].Enabled(ctx, r.Level) {
			continue
		}
		errs = append(errs, h[i].Handle(ctx, r))
	}
	return errors.Join(errs...)
}

// WithAttrs implements the method of the slog.Handler interface
// by cloning the current handler and calling WithAttrs for each
// underlying handler.
func (h handler) WithAttrs(attr []slog.Attr) slog.Handler {
	var nh handler = make(handler, len(h))
	for i := range h {
		nh[i] = h[i].WithAttrs(attr)
	}
	return nh
}

// WithGroup implements the method of the slog.Handler interface
// by cloning the current handler and calling WithGroup for each
// underlying handler.
func (h handler) WithGroup(name string) slog.Handler {
	var nh handler = make(handler, len(h))
	for i := range h {
		nh[i] = h[i].WithGroup(name)
	}
	return nh
}
