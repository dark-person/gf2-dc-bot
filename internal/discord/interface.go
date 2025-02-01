package discord

// Interface for discord bot behaviour.
type Bot interface {
	Notify(message string) error
}
