package messenger

// MessageType represents the type of message being sent.
// It can be either a chat message or a channel message.
type MessageType string

const (
	// MessageTypeChat represents a direct chat message between users.
	MessageTypeChat MessageType = "chat"
	// MessageTypeChannel represents a message sent to a channel within a team.
	MessageTypeChannel MessageType = "channel"
)

// TeamsMessage represents a message to be sent, either to a chat or a channel.
type TeamsMessage struct {
	Type    MessageType
	Content string

	// Channel message fields
	TeamID    string
	ChannelID string

	// Chat message fields
	ChatID string
}
