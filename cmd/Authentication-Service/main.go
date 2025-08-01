package main

import (
	"AythService/internal/config"
	"flag"
	"fmt"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "set-config", "", "Path to configuration file")
	flag.Parse()
	if configPath == "" {
		panic("-set-config flag is required")
	}
	cfg := config.MustLoadConfig(configPath)
	fmt.Printf("Config loaded: %+v\n", cfg)
}
