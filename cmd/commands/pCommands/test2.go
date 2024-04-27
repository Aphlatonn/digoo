package pCommands

import (
	"Gooo/cmd/handler"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func init() {
	fmt.Println("[INFO]: command loaded - test2")
	handler.RegisterCommand(handler.PCommand{
		Name:        "test2",
		Description: "Test command",
		Permissions: discordgo.PermissionAdministrator,
		OwnerOnly:   false,
		Run: func(session *discordgo.Session, message *discordgo.Message, args []string) {
			// Send the message to the channel
			session.ChannelMessageSend(message.ChannelID, "test2 command")
		},
	})
}
