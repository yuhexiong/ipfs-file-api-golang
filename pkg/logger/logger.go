package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(debug bool) {
	var handler slog.Handler

	opts := slog.HandlerOptions{
		AddSource: true,
	}

	if debug {
		opts.Level = slog.LevelDebug
		handler = slog.NewTextHandler(os.Stderr, &opts)
	} else {
		opts.Level = slog.LevelError
		handler = slog.NewJSONHandler(os.Stderr, &opts)
	}

	slog.SetDefault(slog.New(handler))
}
