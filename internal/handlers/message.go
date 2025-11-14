package handlers

import (
	"discord-macro-bot/internal/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MessageHandler handles Discord message events
type MessageHandler struct {
	macroCommand *commands.MacroCommand
}

// NewMessageHandler creates a new message handler
func NewMessageHandler(macroCommand *commands.MacroCommand) *MessageHandler {
	return &MessageHandler{
		macroCommand: macroCommand,
	}
}

// Handle processes incoming Discord messages
func (h *MessageHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore bot messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.ToLower(m.Content)

	// Handle commands
	if strings.HasPrefix(m.Content, "!macro") {
		response, err := h.macroCommand.Handle(m.Content)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "An error occurred processing your request.")
			return
		}
		s.ChannelMessageSendComplex(m.ChannelID, response)
		return
	}

	// Handle keyword triggers (simple responses)
	switch {
	case strings.Contains.tolowercase(content, "nutrition"):
		s.ChannelMessageSend(m.ChannelID, "I can help you find nutrition facts! Use `!macro <food>` to get started.")
	case strings.Contains(content, "bot") && !strings.Contains(content, "robot") && !strings.Contains(content, "bottle"):
		// More specific matching to avoid false positives
		s.ChannelMessageSend(m.ChannelID, "Yes, I am here!")
	}
}

