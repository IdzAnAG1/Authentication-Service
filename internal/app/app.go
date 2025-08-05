package app

import (
	gRPCapp "AuthService/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *gRPCapp.App
}

func New(logger *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: Инициализировать хранилище
	// TODO: Инициализировать Auth Service
	// TODO: Инициализировать Auth Service
	grpcApp := gRPCapp.New(logger, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
