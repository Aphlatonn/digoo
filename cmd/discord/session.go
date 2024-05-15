package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func discord(token string) (*discordgo.Session, error) {
	// Create a new Discord session (session is like client in discord js)
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return nil, err
	}

	// Open a websocket connection to Discord and begin listening for events
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return nil, err
	}

	return dg, nil
}

// get the session
var Session, err = discord("")
