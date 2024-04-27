package pCommands

import (
	"Gooo/cmd/handler"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func init() {
	fmt.Println("[INFO]: command loaded - ping")
	handler.RegisterCommand(handler.PCommand{
		Name:        "test",
		Description: "Test command",
		Run: func(session *discordgo.Session, message *discordgo.Message, args []string) {
			// Create a new embed message
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
				Timestamp: time.Now().Format(time.RFC3339),

				Color: 0x00ff00,
			}

			// Send the embed message to the channel
			session.ChannelMessageSendEmbed(message.ChannelID, embed)
		},
	})
}
