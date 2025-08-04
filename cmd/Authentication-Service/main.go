package main

import (
	"AythService/internal/config"
	"flag"
	"fmt"
	"os"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "set-config", "", "Path to configuration file")
	flag.Parse()
	if configPath == "" {
		panic("-set-config flag is required")
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Error defining the config : %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Config loaded: %+v\n", cfg)
}
