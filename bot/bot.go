package bot



import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)



// store bot api tokens
var (
	OpenNutritionixToken string
	BotToken          string
)



func Run() { 
	// create new discord session

	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	// add event handler for general messages
	discord.AddHandler(newMessage)


	// open session
	discord.Open()
	defer discord.Close()


	// run until code is terminated

	fmt.Println("Bot in running...")
	c := make(chan os.Signal, 1) // channel to recieve os signals to variable c
	signal.Notify(c, os.Interrupt)	// notify channel on interrupt signal
	<-c	// block until a signal is recieved



}



func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// ignore bot messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// switch to respond to messages
	switch{
		case strings.Contains(message.Content, "nutrition"):
			discord.ChannelMessageSend(message.ChannelID, "I can help to find the nutrition facts!")
		case strings.Contains(message.Content, "bot"):
			discord.ChannelMessageSend(message.ChannelID, "Yes, I am here!")

		case strings.Contains(message.Content, "!macro"):
			currentMacro := getNutrition(message.Content)
			discord.ChannelMessageSendComplex(message.ChannelID, currentMacro)
			

	}








}