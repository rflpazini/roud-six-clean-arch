package main

import (
	"log"
	"rflpazini/round-six/config"
	"rflpazini/round-six/internal/server"
)

func main() {
	log.Print("🚀 Server starting...")

	configPath := GetConfigPath("config")
	log.Print(configPath)
	configFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config error: %v", err)
	}

	cfg, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Parsing config error: %v", err)
	}

	s := server.NewServer(cfg)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

// GetConfigPath for local or docker env
func GetConfigPath(config string) string {
	if config == "docker" {
		return "./config/config-docker"
	}

	return "./config/config-local"
}
