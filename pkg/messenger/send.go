package messenger

import (
	"context"
	"fmt"
)

func Send(ctx context.Context, m Messenger, msg Message) error {
	switch msg.Type {
	case MessageTypeChat:
		return m.SendChatMessage(ctx, msg.ChatID, msg)

	case MessageTypeChannel:
		return m.SendChannelMessage(ctx, msg.TeamID, msg.ChannelID, msg)

	default:
		return fmt.Errorf("unsupported message type: %s", msg.Type)
	}
}
