package utils

import (
	"log/slog"
	"os"
)

var (
	Log      *slog.Logger
	LogLevel *slog.LevelVar
)

func init() {
	LogLevel = &slog.LevelVar{}
	opts := slog.HandlerOptions{
		Level: LogLevel,
	}
	handler := slog.NewTextHandler(os.Stdout, &opts)

	Log = slog.New(handler)
}
