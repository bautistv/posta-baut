//go:generate mockgen -source=messenger.go -destination=mocks/messenger.go -package=mocks
package messenger

import (
	"context"
)

// Messenger defines methods for sending messages to chats and channels.
type Messenger interface {
	SendChannelMessage(
		ctx context.Context,
		teamID string,
		channelID string,
		msg string,
	) error
	SendChatMessage(
		ctx context.Context,
		chatID string,
		msg string,
	) error
}
