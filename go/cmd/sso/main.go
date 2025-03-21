package main

import (
	"errors"
	"go-auth-service/internal/config"
	"go-auth-service/internal/lib/logger/handlers/slogpretty"
	"go-auth-service/internal/lib/logger/sl"
	"log/slog"
	"os"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	config := config.MustLoad()

	log := setupLogger(config.Env)

	log.Info("starting application",
		slog.String("env", config.Env),
		slog.Any("grpc", config.GRPC),
		slog.Duration("token_ttl", config.TokenTTL),
	)

	err := errors.New("some error")
	log.Error("error message", sl.Err(err))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}),
		)
	default:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
