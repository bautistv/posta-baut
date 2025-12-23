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
		msg Message,
	) error

	SendChatMessage(
		ctx context.Context,
		chatID string,
		msg Message,
	) error
}
