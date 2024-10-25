package tlog

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

type Logger struct {
	loggers []*slog.Logger
	h       *Handler
	mu      sync.Mutex
}

func New(h *Handler) *Logger {

	loggers := make([]*slog.Logger, len(h.w))

	level := PrepareLevel(h.level)

	for i, writer := range h.w {
		loggers[i] = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: level,
		}))
	}

	return &Logger{
		loggers: loggers,
		h:       h,
	}
}

func (l *Logger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, writer := range l.h.w {
		if file, ok := writer.(*os.File); ok {
			file.Close()
		}
	}

	return
}

func (l *Logger) logWithLevel(ctx context.Context, msg string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, logger := range l.loggers {
		switch l.h.level {
		case LEVEL_INFO:
			logger.InfoContext(ctx, msg, args...)
		case LEVEL_WARN:
			logger.WarnContext(ctx, msg, args...)
		case LEVEL_DEBUG:
			logger.DebugContext(ctx, msg, args...)
		case LEVEL_ERROR:
			logger.ErrorContext(ctx, msg, args...)
		}
	}
}

// Info logs an informational message.
func (l *Logger) Info(ctx context.Context, msg string, args ...any) {
	l.logWithLevel(ctx, msg, args...)
}

// Warn logs a warning message.
func (l *Logger) Warn(ctx context.Context, msg string, args ...any) {
	l.logWithLevel(ctx, msg, args...)
}

// Debug logs a debug message.
func (l *Logger) Debug(ctx context.Context, msg string, args ...any) {
	l.logWithLevel(ctx, msg, args...)
}

// Error logs an error message.
func (l *Logger) Error(ctx context.Context, msg string, args ...any) {
	l.logWithLevel(ctx, msg, args...)
}
