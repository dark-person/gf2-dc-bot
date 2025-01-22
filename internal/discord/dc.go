package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dark-person/gf2-dc-bot/internal/config"
)

// Send message to discord channel. Please note that discord token must be set before call this function.
func SendDiscordMsg(cfg *config.DiscordConfig, message string) error {
	if cfg.Token == "" || cfg.ChannelID == "" {
		return fmt.Errorf("discord token or channel ID not set")
	}

	discordBot, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return fmt.Errorf("failed to create discord bot: %v", err)
	}

	err = discordBot.Open()
	if err != nil {
		return fmt.Errorf("failed to open discord connection: %v", err)
	}

	_, err = discordBot.ChannelMessageSend(cfg.ChannelID, message)
	if err != nil {
		return fmt.Errorf("failed to send message to discord: %v", err)
	}

	return nil
}
