package pCommands

import (
	"Gooo/cmd/handler"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func init() {
	fmt.Println("[INFO]: command loaded - test2")
	handler.RegisterPCommand(handler.PCommand{
		Name:        "test2",
		Description: "Test command",
		Permissions: discordgo.PermissionAdministrator,
		OwnerOnly:   false,
		Run: func(session *discordgo.Session, message *discordgo.Message, args []string) {
			// send an embed with a button to the message channel
			embed := &discordgo.MessageEmbed{
				Title:       "Test Command",
				Description: "testing the descrption baby",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "test field",
						Value: "- test field value",
					},
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: message.Author.AvatarURL("1024"),
				},

				Color: 0x00ff00,
			}

			// create the button
			button := &discordgo.Button{
				Label:    "test button",
				Style:    discordgo.PrimaryButton,
				Disabled: false,
				CustomID: "botona",
			}

			// create the message
			session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
				Content: "test2 command",
				Embeds:  []*discordgo.MessageEmbed{embed},
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							button,
						},
					},
				},
			})
		},
	})
}
