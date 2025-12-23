package svc

import (
	"fmt"

	pb "github.com/bautistv/posta-baut/internal/pb/v1"
	"github.com/bautistv/posta-baut/pkg/messenger"
)

// ReqToMsg converts a SendMessageRequest to a domain-level Message.
func ReqToMsg(req *pb.SendMessageRequest) (messenger.Message, error) {
	var msg messenger.Message

	switch req.MessageType.(type) {
	case *pb.SendMessageRequest_ChannelMessage:
		channelMsg := req.GetChannelMessage()
		msg = messenger.Message{
			Type:      messenger.MessageTypeChannel,
			Content:   channelMsg.GetContent(),
			TeamID:    channelMsg.GetTeamId(),
			ChannelID: channelMsg.GetChannelId(),
		}
	case *pb.SendMessageRequest_ChatMessage:
		chatMsg := req.GetChatMessage()
		msg = messenger.Message{
			Type:    messenger.MessageTypeChat,
			Content: chatMsg.GetContent(),
			ChatID:  chatMsg.GetChatId(),
		}
	default:
		return msg, fmt.Errorf("unsupported message type: %v", req.GetMessageType())
	}

	return msg, nil
}
