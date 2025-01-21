package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Token of discord bot. This token is required to send message to discord channel.
var discordToken string

// Send message to discord channel. Please note that discord token must be set before call this function.
func sendDiscordMsg(channelID string, message string) error {
	if discordToken == "" || channelID == "" {
		return fmt.Errorf("discord token or channel ID not set")
	}

	discordBot, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		return fmt.Errorf("failed to create discord bot: %v", err)
	}

	err = discordBot.Open()
	if err != nil {
		return fmt.Errorf("failed to open discord connection: %v", err)
	}

	_, err = discordBot.ChannelMessageSend(channelID, message)
	if err != nil {
		return fmt.Errorf("failed to send message to discord: %v", err)
	}

	return nil
}
