package handler

import "github.com/bwmarrin/discordgo"

// Command struct
type PCommand struct {
	Name        string
	Description string
	Permissions int64
	OwnerOnly   bool
	Run         func(session *discordgo.Session, message *discordgo.Message, args []string)
}

// Command map
var PCommandMap = map[string]PCommand{}

// RegisterCommand registers a command in the CommandMap
func RegisterPCommand(command PCommand) {
	PCommandMap[command.Name] = command
}
