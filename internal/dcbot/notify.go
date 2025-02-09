package dcbot

import "fmt"

// Send message to discord channel. Please note that discord token must be set before call this function.
func (bm *BotManager) SendReminder(message string) error {
	if !bm.initialized {
		return fmt.Errorf("bot manager not initialized")
	}

	_, err := bm.session.ChannelMessageSend(bm.ReminderChannel, message)
	if err != nil {
		return fmt.Errorf("failed to send message to discord: %v", err)
	}
	return nil
}
