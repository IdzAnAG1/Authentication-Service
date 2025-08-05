package gRPCapp

import (
	authgRPC "AuthService/internal/grpc/auth"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(logger *slog.Logger, port int) *App {
	gRPCServ := grpc.NewServer()
	authgRPC.Register(gRPCServ)
	return &App{
		logger:     logger,
		gRPCServer: gRPCServ,
		port:       port,
	}
}
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
func (a *App) Run() error {
	const op = "gRPCapp.Run"
	log := a.logger.With(
		slog.String("op", op),
		slog.Int("port", a.port))

	log.Info("Starting gRPC Server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	log.Info("gRPC Server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "gRPCapp.Stop"
	a.logger.With(
		slog.String("op", op)).Info("Stopping gRPC Server",
		slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
