package discord

import "fmt"

// Send message to discord channel. Please note that discord token must be set before call this function.
func (bm *BotManager) Notify(message string) error {
	if !bm.initalized {
		return fmt.Errorf("bot manager not initialized")
	}

	_, err := bm.Notifiy.ChannelMessageSend(bm.NotifiyChannel, message)
	if err != nil {
		return fmt.Errorf("failed to send message to discord: %v", err)
	}
	return nil
}
