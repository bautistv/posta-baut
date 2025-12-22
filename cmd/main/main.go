package main

import (
	"log"
	"os"

	"github.com/bautistv/posta-baut/cmd/client"
	config "github.com/bautistv/posta-baut/cmd/config"

	"github.com/bautistv/posta-baut/cmd/server"
	yaml "github.com/goccy/go-yaml"
)

func loadConfig(path string) (config.ClientConfig, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// 3. Declare a variable of your struct type to hold the data
	var config config.ClientConfig

	// 4. Unmarshal the byte slice into the struct
	if err := yaml.Unmarshal(fileBytes, &config); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
	return config, nil
}

func main() {
	cfg, err := loadConfig("./config/local.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
		return
	}

	// Create client configuration
	cli, err := client.NewClient(cfg.MessengerConfig, cfg.LookupClientConfig)
	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
		return
	}

	// Create and run the server
	server, err := server.NewServer(cli)
	if err != nil {
		log.Fatalf("Failed to create server: %v\n", err)
		return
	}
	server.Run()
}
