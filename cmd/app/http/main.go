package main

import (
	"fmt"
	"log"

	"github.com/api-sekejap/config"
)

const (
	configPath = "../../config/manager"
)

func main() {
	config, err := config.NewConfigManager(configPath)
	if err != nil {
		log.Fatalf("Error when loading config file %v", err)
		return
	}

	fmt.Println(config)
}
