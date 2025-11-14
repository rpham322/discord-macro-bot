package commands

import (
	"discord-macro-bot/internal/api"
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const macroCommandPrefix = "!macro"

// MacroCommand handles the !macro command
type MacroCommand struct {
	apiClient *api.Client
}

// NewMacroCommand creates a new macro command handler
func NewMacroCommand(apiClient *api.Client) *MacroCommand {
	return &MacroCommand{
		apiClient: apiClient,
	}
}

// Handle processes the macro command and returns a Discord message
func (c *MacroCommand) Handle(message string) (*discordgo.MessageSend, error) {
	// Parse command
	query, err := c.parseCommand(message)
	if err != nil {
		return &discordgo.MessageSend{
			Content: fmt.Sprintf("Error: %s\nUsage: `%s <food>`", err.Error(), macroCommandPrefix),
		}, nil
	}

	// Get nutrition data
	food, err := c.apiClient.GetNutrition(query)
	if err != nil {
		return &discordgo.MessageSend{
			Content: fmt.Sprintf("Could not find nutrition data for **%s**\n%s", query, err.Error()),
		}, nil
	}

	// Build embed response
	return c.buildEmbed(food), nil
}

// parseCommand extracts the food query from the command
func (c *MacroCommand) parseCommand(message string) (string, error) {
	r := regexp.MustCompile(`(?i)^` + regexp.QuoteMeta(macroCommandPrefix) + `\s+(.+)$`)
	m := r.FindStringSubmatch(strings.TrimSpace(message))

	if len(m) < 2 {
		return "", fmt.Errorf("Please provide a food item after the command")
	}

	return m[1], nil
}

// buildEmbed creates a Discord embed from food data
func (c *MacroCommand) buildEmbed(food *api.Food) *discordgo.MessageSend {
	name := strings.Title(food.FoodName)
	brand := ""
	if food.BrandName != nil && *food.BrandName != "" {
		brand = " (" + *food.BrandName + ")"
	}

	serving := fmt.Sprintf("%.0f %s (%.0fg)", food.ServingQty, food.ServingUnit, food.ServingWeightGrams)
	calories := fmt.Sprintf("%.0f kcal", food.Calories)
	protein := fmt.Sprintf("%.1f g", food.Protein)
	carbs := fmt.Sprintf("%.1f g", food.Carbs)
	fat := fmt.Sprintf("%.1f g", food.Fat)
	fiber := fmt.Sprintf("%.1f g", food.Fiber)
	sugars := fmt.Sprintf("%.1f g", food.Sugars)

	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Type:        discordgo.EmbedTypeRich,
				Title:       "Nutrition — " + name + brand,
				Description: "Serving: " + serving,
				Fields: []*discordgo.MessageEmbedField{
					{Name: "Calories", Value: calories, Inline: true},
					{Name: "Protein", Value: protein, Inline: true},
					{Name: "Carbs", Value: carbs, Inline: true},
					{Name: "Fat", Value: fat, Inline: true},
					{Name: "Fiber", Value: fiber, Inline: true},
					{Name: "Sugars", Value: sugars, Inline: true},
				},
			},
		},
	}
}

