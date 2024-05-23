package logger

import (
	"log/slog"
	"os"
	"strings"
)

func SetupGlobalLogger(config *LoggerConfig) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)
}

func mapLoggerLevel(level string) slog.Level {
	level = strings.ToUpper(level)
	switch level {
	case slog.LevelInfo.Level().String():
		return slog.LevelInfo
	case slog.LevelError.Level().String():
		return slog.LevelError
	case slog.LevelWarn.Level().String():
		return slog.LevelWarn
	case slog.LevelDebug.Level().String():
		return slog.LevelDebug
	}
	return slog.LevelInfo
}
