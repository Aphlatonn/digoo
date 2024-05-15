package utils

import (
	"Digoo/cmd/handler"
	"github.com/bwmarrin/discordgo"
)

func init() {
	handler.PrefixCommandBuilder().
		SetName("ping").
		SetDescription("Ping command to test responsiveness").
		SetPermissions(discordgo.PermissionSendMessages).
		SetOwnerOnly(false).
		SetRunFunction(func(session *discordgo.Session, message *discordgo.Message, args []string) {
			session.ChannelMessageSend(message.ChannelID, "Pong!")
		}).
		Build()
}
