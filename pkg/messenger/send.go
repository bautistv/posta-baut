package messenger

import (
	"context"
	"fmt"
)

// Send sends a message using the provided Messenger based on the MessageType.
func Send(ctx context.Context, m Messenger, msg TeamsMessage) error {
	if m == nil {
		return fmt.Errorf("messenger cannot be nil")
	}

	switch msg.Type {
	case MessageTypeChat:
		if msg.ChatID == "" {
			return fmt.Errorf("missing chat ID from chat message")
		}
		if msg.Content == "" {
			return fmt.Errorf("message content cannot be empty")
		}
		return m.SendChatMessage(ctx, msg.ChatID, msg)

	case MessageTypeChannel:
		if msg.TeamID == "" {
			return fmt.Errorf("missing team ID from channel message")
		}
		if msg.ChannelID == "" {
			return fmt.Errorf("missing channel ID from channel message")
		}
		if msg.Content == "" {
			return fmt.Errorf("message content cannot be empty")
		}
		return m.SendChannelMessage(ctx, msg.TeamID, msg.ChannelID, msg)

	default:
		return fmt.Errorf("unsupported message type: %s", msg.Type)
	}
}
