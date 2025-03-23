package grpcapp

import (
	"fmt"
	authgrpc "go-auth-service/internal/grpc/auth"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	authService authgrpc.AuthService,
	port int,
) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer, authService)

	return &App{log: log, port: port, gRPCServer: gRPCServer}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op))

	log.Info("starting grpc server", slog.Int("port", a.port))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", listener.Addr().String()))

	if err := a.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping grpc server")

	defer a.gRPCServer.GracefulStop()
}
