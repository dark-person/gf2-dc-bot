// Package for common interface.
package api

import "github.com/bwmarrin/discordgo"

// Interface for discord bot behaviour.
type DiscordBot interface {
	SendReminder(message string) error
}

// Persona interface for discord bots.
//
// This interface is defined for easily implement for other bots.
type BotPersona interface {
	// Return help message
	Help() string

	// Return random character dialog as string.
	RandomDialog() string

	// Send message if message mention current bot as user.
	ReplyIfMentionBot(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error)

	// Send message if message mention current bot name.
	ReplyIfMentionBotName(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error)

	// Send message if message contains keyword to trigger special reaction.
	ReplyIfHasKeyword(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error)
}
