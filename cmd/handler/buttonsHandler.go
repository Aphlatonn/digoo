package handler

import "github.com/bwmarrin/discordgo"

// intraction struct
type Button struct {
	Id          string
	Permissions int64
	OwnerOnly   bool
	Run         func(session *discordgo.Session, intraction *discordgo.InteractionCreate)
}

// buttons map
var ButtonsMap = map[string]Button{}

// RegisterButton registers a button in the buttonsMap
func RegisterButton(button Button) {
	ButtonsMap[button.Id] = button
}
