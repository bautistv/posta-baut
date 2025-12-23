package messenger

import (
	"context"
)

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
