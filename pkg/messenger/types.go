package messenger

type MessageType string

const (
	MessageTypeChat    MessageType = "chat"
	MessageTypeChannel MessageType = "channel"
)

type Message struct {
	Type    MessageType
	Content string

	// Channel message fields
	TeamID    string
	ChannelID string

	// Chat message fields
	ChatID string
}
