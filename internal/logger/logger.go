package logger

import (
	"log/slog"
	"os"
)

const (
	Local = "local"
	Dev   = "dev"
	Prod  = "prod"
)

func SetupLogger(LogLevel string) *slog.Logger {
	var log *slog.Logger

	switch LogLevel {
	case Local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case Dev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case Prod:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	return log
}
