package multislog

import (
	"context"
	"errors"
	"log/slog"
)

type handler []slog.Handler

func NewHandler(handlers ...slog.Handler) slog.Handler {
	return handler(handlers)
}

// Enabled implements the method of the slog.Handler interface
// by calling the same method of the formatter habdler.
func (h handler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

// Handle implements the method of the slog.Handler interface.
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

// WithAttrs implements the method of the slog.Handler interface by
// cloning the current handler and calling the WithAttrs of the
// formatter handler.
func (h handler) WithAttrs(attr []slog.Attr) slog.Handler {
	var nh handler = make(handler, len(h))
	for i := range h {
		nh[i] = h[i].WithAttrs(attr)
	}
	return nh
}

// WithGroup implements the method of the slog.Handler interface by
// cloning the current handler and calling the WithGroup of the
// formatter handler.
func (h handler) WithGroup(name string) slog.Handler {
	var nh handler = make(handler, len(h))
	for i := range h {
		nh[i] = h[i].WithGroup(name)
	}
	return nh
}
