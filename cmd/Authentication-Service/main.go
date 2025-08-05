package main

import (
	"AuthService/internal/app"
	"AuthService/internal/config"
	"AuthService/internal/logger"
	"fmt"
	"os"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error defining the config : %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Config loaded: %+v\n", cfg)
	log := logger.SetupLogger(cfg.Logger.Level)
	log.Info("Logger is Activate")

	application := app.New(
		log,
		cfg.GRPCServer.Port,
		cfg.Storage.Path,
		cfg.GRPCServer.Interval)

	application.GRPCSrv.MustRun()
}
