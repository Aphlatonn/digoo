package events

import (
	"fmt"
	"strings"

	"Gooo/cmd/handler"

	"github.com/bwmarrin/discordgo"
)

const prefix = "!!"

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	// check if the message author is the same as the bot
	if message.Author.ID == session.State.User.ID {
		return
	}

	// check if the message is from a bot
	if message.Author.Bot {
		return
	}

	// Check if the message starts with the prefix
	if !strings.HasPrefix(message.Content, prefix) {
		return

	}

	// Split the message content into command and arguments
	parts := strings.Fields(message.Content)
	// Remove the prefix from the command name
	commandName := parts[0][len(prefix):]
	args := parts[1:]

	// Check if the command exists
	cmd, ok := handler.PCommandMap[commandName]
	if !ok {
		// Command not found
		return
	}

	fmt.Println(cmd.Permissions)

	// Check for the user permisession.LastHeartbeatAck.Unixns
	if cmd.Permissions != 0 {
		p, e := session.State.MessagePermissions(message.Message)

		fmt.Println(p & cmd.Permissions)

		// catch errors
		if e != nil {
			session.ChannelMessageSend(message.ChannelID, "An error occurred while checking your permission")
			fmt.Println(e.Error())
			return
		}

		if p&cmd.Permissions != cmd.Permissions {
			session.ChannelMessageSend(message.ChannelID, "You do not have permission to use this command")
			return
		}
	}

	// Check if the user is message guild owner
	if cmd.OwnerOnly {
		//get the owner id
		Guild, e := session.Guild(message.GuildID)

		// check for errors
		if e != nil {
			session.ChannelMessageSend(message.ChannelID, "An error occurred while checking if you are the guild owner")
			fmt.Println(e.Error())
			return
		}

		if message.Author.ID != Guild.OwnerID {
			session.ChannelMessageSend(message.ChannelID, "You do not have permission to use this command")
			return
		}
	}

	// Execute the command's Run function with the necessary parameters
	cmd.Run(session, message.Message, args)
}
