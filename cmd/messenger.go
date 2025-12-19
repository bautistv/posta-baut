package posta_baut

import (
	"context"
)

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
