package bot

import (
	"discord-macro-bot/internal/api"
	"discord-macro-bot/internal/commands"
	"discord-macro-bot/internal/config"
	"discord-macro-bot/internal/handlers"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Bot represents the Discord bot instance
type Bot struct {
	config  *config.Config
	session *discordgo.Session
}

// New creates a new bot instance
func New(cfg *config.Config) *Bot {
	return &Bot{
		config: cfg,
	}
}

// Start initializes and starts the Discord bot
func (b *Bot) Start() error {
	// Create Discord session
	discord, err := discordgo.New("Bot " + b.config.DiscordToken)
	if err != nil {
		return fmt.Errorf("Failed to create Discord session: %w", err)
	}

	b.session = discord

	// Create API client
	apiClient := api.NewClient(b.config.NutritionixAppID, b.config.NutritionixToken)

	// Create command handlers
	macroCommand := commands.NewMacroCommand(apiClient)

	// Create message handler
	messageHandler := handlers.NewMessageHandler(macroCommand)

	// Register message handler
	discord.AddHandler(messageHandler.Handle)

	// Open Discord session
	if err := discord.Open(); err != nil {
		return fmt.Errorf("Failed to open Discord session: %w", err)
	}

	log.Println("Bot is running... Press CTRL-C to exit.")

	// Wait for interrupt signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Bot is shutting down...")
	
	// Close Discord session cleanly
	if err := discord.Close(); err != nil {
		log.Printf("Error closing Discord session: %v", err)
	}
	
	return nil
}

