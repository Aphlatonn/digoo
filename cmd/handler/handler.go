package handler

import "github.com/bwmarrin/discordgo"

// ////////////////////////////////////////////////////////////// prefix commands things
// Command represents a command
type PCommand struct {
	Name        string
	Description string
	Permissions int64
	OwnerOnly   bool
	Run         func(session *discordgo.Session, message *discordgo.Message, args []string)
}

// CommandMap is a map of command names to commands
var PCommandMap = map[string]PCommand{}

// RegisterCommand registers a command in the CommandMap
func RegisterCommand(command PCommand) {
	PCommandMap[command.Name] = command
}
