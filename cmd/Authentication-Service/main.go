package main

import (
	"AythService/internal/config"
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
}
