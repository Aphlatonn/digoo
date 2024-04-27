package events

import (
	"Gooo/cmd/handler"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func ButtonClick(session *discordgo.Session, intraction *discordgo.InteractionCreate) {
	ID := intraction.MessageComponentData().CustomID

	button, ok := handler.ButtonsMap[ID]
	if !ok {
		// button not found
		return
	}

	// Check if the user is message guild owner
	if button.OwnerOnly {
		//get the owner id
		Guild, e := session.Guild(intraction.GuildID)

		// check for errors
		if e != nil {
			session.ChannelMessageSend(intraction.ChannelID, "An error occurred while checking if you are the guild owner")
			fmt.Println(e.Error())
			return
		}

		if intraction.User.ID != Guild.OwnerID {
			session.ChannelMessageSend(intraction.ChannelID, "You do not have permission to use this command")
			return
		}
	}

	// Execute the command's Run function with the necessary parameters
	button.Run(session, intraction)
}
