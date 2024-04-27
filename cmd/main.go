package main

import (
	_ "Gooo/cmd/commands/pCommands"
	"Gooo/cmd/events"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// ======================== Register Events ========================//
	// Register the Ready event handler
	dg.AddHandler(events.Ready)

	// Register the MessageCreate event handler
	dg.AddHandler(events.MessageCreate)
	// ================================================================//

	// Open a websocket connection to Discord and begin listening for events
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	// Keep the bot running indefinitely
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	<-make(chan struct{})
}
