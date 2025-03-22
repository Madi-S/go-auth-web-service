package main

import (
	"go-auth-service/internal/app"
	"go-auth-service/internal/config"
	"go-auth-service/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	config := config.MustLoad()

	log := setupLogger(config.Env)

	log.Info("starting application", slog.Any("config", config))

	application := app.New(log, config.GRPC.Port, config.StoragePath, config.TokenTTL)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGABRT)

	sig := <-stop
	log.Info("received signal", slog.String("signal", sig.String()))

	application.GRPCServer.Stop()
	log.Info("application stopped")
	wg.Wait()
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
