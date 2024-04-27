package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Ready is a function to handle the "ready" event
func Ready(session *discordgo.Session, event *discordgo.Ready) {
	// Ready event handler logic
	fmt.Println(session.State.User.Username + " Bot is ready!")
}
