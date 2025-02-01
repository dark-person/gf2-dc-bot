package discord

// Interface for discord bot behaviour.
type Bot interface {
	SendReminder(message string) error
}
