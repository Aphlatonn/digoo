package client

import (
	"Digoo/cmd/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Ready event
func init() {
	fmt.Println("[INFO] ready event Registered")
	discord.Session.AddHandler(func(session *discordgo.Session, event *discordgo.Ready) {
		fmt.Println("[INFO] " + session.State.User.Username + " Bot is ready!")
	})
}
