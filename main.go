package main



import (	
	"discord-macro-bot/bot"
	"log"
	"os"


	"github.com/joho/godotenv"

)


func main() {



	// Load .env falls back to OS env if file missing
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}


	// Load Discord bot token
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("DISCORD_BOT_TOKEN is not set")
	}

	// Load Nutritionix tokens
	nutritionix_app_id := os.Getenv("NUTRITIONIX_APP_ID")
	nutritionix_token:= os.Getenv("NUTRITIONIX_TOKEN") 
	if nutritionix_app_id == "" || nutritionix_token == "" {
		log.Fatal("NUTRITIONIX_APP_ID and/or NUTRITIONIX_TOKEN are not set")
	}

	// Start the bot
	bot.BotToken = botToken
	bot.OpenNutritionixToken = nutritionix_token
	bot.Run()











}