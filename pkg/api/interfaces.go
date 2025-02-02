// Package for common interface.
package api

// Interface for discord bot behaviour.
type DiscordBot interface {
	SendReminder(message string) error
}
