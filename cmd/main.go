package main

import (
	"discord-macro-bot/internal/bot"
	"discord-macro-bot/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (falls back to OS env if file missing)
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create and start bot
	discordBot := bot.New(cfg)
	if err := discordBot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}

