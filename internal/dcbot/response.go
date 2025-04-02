package dcbot

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

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

	fmt.Printf("%s [%s] %s: %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		m.ChannelID, m.Author.Username, m.Content)

	var isSent bool
	var err error

	// help command
	if m.Content == "?help" {
		err := discordutils.SendMsgToChannel(s, m.ChannelID, bm.persona.Help())
		if err != nil {
			fmt.Println(err)
			return
		}

		return // Early return to prevent duplicate messages
	}

	// ----------------------------------------

	// Check if reminder command calls
	if m.Content == "!remind" {
		err := discordutils.SendMsgToChannel(s, m.ChannelID, bm.getDailyReminderMsg())
		if err != nil {
			fmt.Println(err)
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
		fmt.Println(err)
		return
	}

	// ----------------------------------------

	// Check if bot is mentioned by '@' method
	isSent, err = bm.persona.ReplyIfMentionBot(s, m.ChannelID, m.Content)
	if isSent {
		return // Early return to prevent duplicate messages
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	// ----------------------------------------

	// Check if special reaction
	isSent, err = bm.persona.ReplyIfHasKeyword(s, m.ChannelID, m.Content)
	if isSent {
		return // Early return to prevent duplicate messages
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}
