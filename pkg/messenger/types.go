package messenger

// TeamsChannelTarget represents a target destination in a Microsoft Teams channel.
// It includes the Team ID, Channel ID, and optionally a Thread ID to specify a thread reply.
type TeamsChannelTarget struct {
	// TeamID is the identifier of the Microsoft Teams team.
	TeamID string

	// ChannelID is the identifier of the channel within the team.
	ChannelID string

	// ThreadID is the optional identifier of the root message of a thread.
	// If set, the message is sent as a reply in the specified thread.
	ThreadID string
}

// ChatTarget represents a target destination in a Microsoft Teams chat.
// It includes the Chat ID and optionally a message ID to reply to within the chat.
type ChatTarget struct {
	// ChatID is the identifier of the chat (1:1 or group chat).
	ChatID string

	// ReplyToMessageID is the optional identifier of the message being replied to.
	// If set, the message is sent as a reply to this message in the chat.
	ReplyToMessageID string
}

// MessageTarget represents the destination of a message.
// It supports either a Teams channel target or a chat target, but not both.
type MessageTarget struct {
	// Channel specifies the Teams channel target.
	Channel *TeamsChannelTarget

	// Chat specifies the Teams chat target.
	Chat *ChatTarget
}

// SendMessageRequest represents a request to send a message to a specified target.
type SendMessageRequest struct {
	// Target specifies where the message should be sent.
	Target MessageTarget

	// Content is the body text of the message to be sent.
	Content string
}

// SendMessageResponse represents the response after sending a message.
type SendMessageResponse struct {
	// MessageID is the identifier of the sent message.
	MessageID string
}
