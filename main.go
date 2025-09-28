package main



import (	
	"discord-macro-bot/bot"
	"log"
	"os"



)


func main() {
	//load env variables
	botToken, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		log.Fatal("Must set DISCORD_BOT_TOKEN environment variable")
	}
	openNutrionixToken, ok := os.LookupEnv("NUTRIONIX_TOKEN")
	if !ok {
		log.Fatal("Must set NUTRIONIX_TOKEN environment variable")
	}

	//start the bot
	bot.BotToken = botToken
	bot.OpenNutrionixToken = openNutrionixToken
	bot.Run()














}