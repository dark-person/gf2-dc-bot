package dcbot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/dark-person/gf2-dc-bot/pkg/discordutils"
)

// This function will be called (due to AddHandler above)
// every time a new message is created on any channel
// that the authenticated bot has access to.
func (bm *BotManager) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself, which is a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Trace().
		Timestamp().
		Str("Channel", m.ChannelID).
		Str("Author", m.Author.Username).
		Str("Msg", m.Content).
		Send()

	var isSent bool
	var err error

	// help command
	if m.Content == "?help" {
		err := discordutils.SendMsgToChannel(s, m.ChannelID, bm.persona.Help())
		if err != nil {
			log.Error().Err(err).Send()
			return
		}

		return // Early return to prevent duplicate messages
	}

	// ----------------------------------------

	// Check if reminder command calls
	if m.Content == "!remind" {
		err := discordutils.SendMsgToChannel(s, m.ChannelID, bm.getDailyReminderMsg())
		if err != nil {
			log.Error().Err(err).Send()
			return
		}

		return // Early return to prevent duplicate messages
	}

	// ----------------------------------------

	// Check if message is mention bot name
	isSent, err = bm.persona.ReplyIfMentionBotName(s, m.ChannelID, m.Content)
	if isSent {
		return // Early return to prevent duplicate messages
	}

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	// ----------------------------------------

	// Check if bot is mentioned by '@' method
	isSent, err = bm.persona.ReplyIfMentionBot(s, m.ChannelID, m.Content)
	if isSent {
		return // Early return to prevent duplicate messages
	}

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	// ----------------------------------------

	// Check if special reaction
	isSent, err = bm.persona.ReplyIfHasKeyword(s, m.ChannelID, m.Content)
	if isSent {
		return // Early return to prevent duplicate messages
	}

	if err != nil {
		log.Error().Err(err).Send()
		return
	}
}
