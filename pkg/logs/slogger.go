package logs

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type Level slog.Level

var (
	LevelDebug Level = Level(slog.LevelDebug)
	LevelInfo  Level = Level(slog.LevelInfo)
	LevelWarn  Level = Level(slog.LevelWarn)
	LevelError Level = Level(slog.LevelError)
	LevelFatal Level = Level(slog.LevelError) // slog does not have Fatal level
)

func (level Level) String() string {
	if level == LevelFatal {
		return "FATAL"
	}

	return slog.Level(level).String()
}

// Slogger is a wrapper around slog.Logger that implements the Logger interface.
type Slogger struct {
	*slog.Logger
}

// NewMultiSlogger returns a new Slogger that can write to multiple handlers.
// i.e. if you want to write to both a file and the console, you can create a MultiSlogger with both handlers.
func NewMultiSlogger(handlers ...slog.Handler) Logger {
	return &Slogger{slog.New(slog.NewMultiHandler(handlers...))}
}

// NewNullSlogger returns a logger that discards all log messages.
// This can be used as a default logger to avoid nil pointer dereference when a logger is not provided.
func NewNullSlogger() Logger {
	return &Slogger{slog.New(slog.NewJSONHandler(io.Discard, nil))}
}

// With returns a new Slogger with the provided key-value pairs added to the output of all log messages.
func (sl *Slogger) With(args ...any) Logger {
	return &Slogger{sl.Logger.With(args...)}
}

// InfoContext logs a message at the Info level with the provided context and key-value pairs.
func (sl *Slogger) InfoContext(ctx context.Context, msg string, args ...any) {
	sl.Logger.InfoContext(ctx, msg, args...)
}

// DebugContext logs a message at the Debug level with the provided context and key-value pairs.
func (sl *Slogger) DebugContext(ctx context.Context, msg string, args ...any) {
	sl.Logger.DebugContext(ctx, msg, args...)
}

// WarnContext logs a message at the Warn level with the provided context and key-value pairs.
func (sl *Slogger) WarnContext(ctx context.Context, msg string, args ...any) {
	sl.Logger.WarnContext(ctx, msg, args...)
}

// ErrorContext logs a message at the Error level with the provided context and key-value pairs.
func (sl *Slogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	sl.Logger.ErrorContext(ctx, msg, args...)
}

// Log logs a message at the specified level with the provided context and key-value pairs.
func (sl *Slogger) Log(ctx context.Context, level Level, msg string, args ...any) {
	sl.Logger.Log(ctx, slog.Level(level), msg, args...)
}

// FatalContext logs a message at the Fatal level with the provided context and key-value pairs, and then exits the application.
func (sl *Slogger) FatalContext(ctx context.Context, msg string, args ...any) {
	sl.Logger.Log(ctx, slog.LevelError, msg, args...)

	// TODO: slog does not have Fatal level, so we exit the application here
	// I'm not certain this gives enough time for the logs on both handlers to be flushed?

	os.Exit(1)
}
