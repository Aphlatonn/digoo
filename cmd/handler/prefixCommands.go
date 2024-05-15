package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// PCommand struct
type PCommand struct {
	Name        string
	Aliases     []string
	Description string
	Permissions int64
	OwnerOnly   bool
	Run         func(session *discordgo.Session, message *discordgo.Message, args []string)
}

// PCommandBuilder struct
type PCommandBuilder struct {
	command PCommand
}

// NewPCommandBuilder creates a new PCommandBuilder instance
func PrefixCommandBuilder() *PCommandBuilder {
	return &PCommandBuilder{command: PCommand{}}
}

// SetName sets the name of the command
func (builder *PCommandBuilder) SetName(name string) *PCommandBuilder {
	builder.command.Name = name
	return builder
}

// SetAliases sets the aliases of the command
func (builder *PCommandBuilder) SetAliases(aliases []string) *PCommandBuilder {
	builder.command.Aliases = aliases
	return builder
}

// SetDescription sets the description of the command
func (builder *PCommandBuilder) SetDescription(description string) *PCommandBuilder {
	builder.command.Description = description
	return builder
}

// SetPermissions sets the permissions required for the command
func (builder *PCommandBuilder) SetPermissions(permissions int64) *PCommandBuilder {
	builder.command.Permissions = permissions
	return builder
}

// SetOwnerOnly sets whether the command is only for the owner
func (builder *PCommandBuilder) SetOwnerOnly(ownerOnly bool) *PCommandBuilder {
	builder.command.OwnerOnly = ownerOnly
	return builder
}

// SetRunFunction sets the function to run when the command is executed
func (builder *PCommandBuilder) SetRunFunction(runFunc func(session *discordgo.Session, message *discordgo.Message, args []string)) *PCommandBuilder {
	builder.command.Run = runFunc
	return builder
}

// Build builds the PCommand object and registers it in the CommandMap
func (builder *PCommandBuilder) Build() PCommand {
	command := builder.command
	RegisterPCommand(command)
	return command
}

// RegisterPCommand registers a command in the CommandMap
func RegisterPCommand(command PCommand) {
	PCommandMap[command.Name] = command
	fmt.Println("[INFO] command Registered: " + command.Name)
}

// Command map
var PCommandMap = map[string]PCommand{}
