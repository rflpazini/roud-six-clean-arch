package main

import (
	"log"
	"rflpazini/round-six/config"
	"rflpazini/round-six/internal/server"

	"go.uber.org/zap"
)

func main() {
	log.Print("🚀 Server starting...")
	configs, err := loadConfigs()
	if err != nil {
		return
	}

	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	s := server.NewServer(configs, logger)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

func loadConfigs() (*config.Config, error) {
	configPath := config.GetConfigPath("config")
	configFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config error: %v", err)
	}

	return config.ParseConfig(configFile)
}
