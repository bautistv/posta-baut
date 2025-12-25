// Package utils provides service-level utilities for the Posta Baut application.
package utils

import (
	"fmt"

	pb "github.com/bautistv/posta-baut/internal/pb/v1"
	"github.com/bautistv/posta-baut/pkg/messenger"
)

// ReqToMsg converts a SendMessageRequest to a domain-level Message.
func ReqToMsg(req *pb.SendMessageRequest) (messenger.TeamsMessage, error) {
	var msg messenger.TeamsMessage

	switch req.MessageType.(type) {
	case *pb.SendMessageRequest_ChannelMessage:
		channelMsg := req.GetChannelMessage()
		msg = messenger.TeamsMessage{
			Type:      messenger.MessageTypeChannel,
			Content:   channelMsg.GetContent(),
			TeamID:    channelMsg.GetTeamId(),
			ChannelID: channelMsg.GetChannelId(),
		}
	case *pb.SendMessageRequest_ChatMessage:
		chatMsg := req.GetChatMessage()
		msg = messenger.TeamsMessage{
			Type:    messenger.MessageTypeChat,
			Content: chatMsg.GetContent(),
			ChatID:  chatMsg.GetChatId(),
		}
	default:
		return msg, fmt.Errorf("unsupported message type: %v", req.GetMessageType())
	}

	return msg, nil
}
