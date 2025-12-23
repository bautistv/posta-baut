// Package svc provides service-level utilities for the Posta Baut application.
package svc

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/bautistv/posta-baut/cmd/client"
	utils "github.com/bautistv/posta-baut/cmd/svc/utils"
	pb "github.com/bautistv/posta-baut/internal/pb/v1"
	"github.com/bautistv/posta-baut/pkg/messenger"
)

// teamsService implements the TeamsService defined in the protobuf.
type teamsService struct {
	Client *client.Client
}

// NewTeamsServiceClient creates a new TeamsService with the provided client.
func NewTeamsServiceClient(client *client.Client) *teamsService {
	return &teamsService{
		Client: client,
	}
}

// SendMessage handles sending messages to Microsoft Teams.
func (s *teamsService) SendMessage(ctx context.Context, req *connect.Request[pb.SendMessageRequest]) (*connect.Response[pb.SendMessageResponse], error) {
	// Convert the request to domain-level Message - messenger.Message
	msg, err := utils.ReqToMsg(req.Msg)
	if err != nil {
		return nil, fmt.Errorf("failed to convert request to message: %w", err)
	}

	// Use the client's Messenger to send the message
	err = messenger.Send(ctx, s.Client.Messenger, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	resp := &pb.SendMessageResponse{
		Success: true,
	}
	return connect.NewResponse(resp), nil
}
