// Package for common interface.
package api

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Interface for discord bot behavior.
type DiscordBot interface {
	SendReminder() error
}

// Persona interface for discord bots.
//
// This interface is defined for easily implement for other bots.
type BotPersona interface {
	// Return help message
	Help() string

	// Get daily reminder message prefix & suffix, for customization.
	GetReminderCustomizedStr(t time.Time) (prefix string, suffix string)

	// Return random character dialog as string.
	RandomDialog() string

	// Send message if message mention current bot as user.
	ReplyIfMentionBot(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error)

	// Send message if message mention current bot name.
	ReplyIfMentionBotName(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error)

	// Send message if message contains keyword to trigger special reaction.
	ReplyIfHasKeyword(s *discordgo.Session, channelID string, incoming string) (isSent bool, err error)
}
