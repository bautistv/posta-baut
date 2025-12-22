package svc

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/bautistv/posta-baut/cmd/client"
	"github.com/bautistv/posta-baut/cmd/messenger"
	"github.com/bautistv/posta-baut/internal/pb"
)

// Your service implementation
type teamsService struct {
	Client *client.Client
}

func (s *teamsService) SendMessage(ctx context.Context, req *connect.Request[pb.SendMessageRequest]) (*connect.Response[pb.SendMessageResponse], error) {
	switch req.Msg.MessageType.(type) {
	case *pb.SendMessageRequest_ChannelMessage:
		// Extract channel message details
		msg := req.Msg.GetChannelMessage()
		channelId := msg.GetChannelId()
		teamId := msg.GetTeamId()
		msgContent := messenger.Message{Content: msg.GetContent()}

		// Send the channel message using the client
		err := s.Client.Messenger.SendChannelMessage(ctx, teamId, channelId, msgContent)
		if err != nil {
			return nil, err
		}
	case *pb.SendMessageRequest_ChatMessage:
		// Extract chat message details
		msg := req.Msg.GetChatMessage()
		chatId := msg.GetChatId()
		msgContent := messenger.Message{Content: msg.GetContent()}

		// Send the chat message using the client
		err := s.Client.Messenger.SendChatMessage(ctx, chatId, msgContent)
		if err != nil {
			return nil, err
		}

	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unsupported message type: %v", req.Msg.GetMessageType()))
	}

	resp := &pb.SendMessageResponse{
		Success: true,
	}
	return connect.NewResponse(resp), nil
}
