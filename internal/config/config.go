package config

import (
	"fmt"
	"os"
)

// Config holds all configuration for the application
type Config struct {
	DiscordToken      string
	NutritionixAppID  string
	NutritionixToken  string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	discordToken := os.Getenv("DISCORD_BOT_TOKEN")
	if discordToken == "" {
		return nil, fmt.Errorf("DISCORD_BOT_TOKEN is not set")
	}

	nutritionixAppID := os.Getenv("NUTRITIONIX_APP_ID")
	if nutritionixAppID == "" {
		return nil, fmt.Errorf("NUTRITIONIX_APP_ID is not set")
	}

	nutritionixToken := os.Getenv("NUTRITIONIX_TOKEN")
	if nutritionixToken == "" {
		return nil, fmt.Errorf("NUTRITIONIX_TOKEN is not set")
	}

	return &Config{
		DiscordToken:     discordToken,
		NutritionixAppID: nutritionixAppID,
		NutritionixToken: nutritionixToken,
	}, nil
}

