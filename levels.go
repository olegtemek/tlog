package tlog

import "log/slog"

type Level string

const (
	LEVEL_INFO  Level = "info"
	LEVEL_ERROR Level = "error"
	LEVEL_WARN  Level = "warn"
	LEVEL_DEBUG Level = "debug"
)

func PrepareLevel(level Level) (slevel slog.Level) {

	switch level {
	case LEVEL_ERROR:
		slevel = slog.LevelError
	case LEVEL_INFO:
		slevel = slog.LevelInfo
	case LEVEL_WARN:
		slevel = slog.LevelWarn
	case LEVEL_DEBUG:
		slevel = slog.LevelDebug
	default:
		slevel = slog.LevelDebug
	}

	return
}
