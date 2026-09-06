// Package for common interface.
package api

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Interface for discord bot behavior.
type DiscordBot interface {
	// Send message to discord channel. Please note that discord token must be set before call this function.
	SendReminder() error

	// Send message to discord channel, to notify specific role to remember gun-smoke frontline event.
	// Please note that discord token must be set before call this function.
	SendGunSmokeReminder() error

	// Send message to discord channel, to notify specific role to remember Intelligence Supplies event.
	// Please note that discord token must be set before call this function.
	SendIntelligenceSuppliesReminder() error

	// Send message to discord channel, to notify specific role to remember Frontier Conquest event.
	// Please note that discord token must be set before call this function.
	SendFrontierConquestReminder() error
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

	// Return message for gun-smoke frontline dialog. The custom string MUST include "<@&%s>" for mention.
	GunSmokeRemindDialog(roleID string) string

	// Return message for claim IntelligenceSupplies dialog. The custom string MUST include "<@&%s>" for mention.
	IntelligenceSuppliesRemindDialog(roleID string) string

	// Return message for Frontier Conquest dialog. The custom string MUST include "<@&%s>" for mention.
	FrontierConquestRemindDialog(roleID string) string
}
