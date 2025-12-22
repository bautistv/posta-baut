package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/bautistv/posta-baut/cmd/client"
	config "github.com/bautistv/posta-baut/cmd/config"
	pb "github.com/bautistv/posta-baut/internal/pb"
	pbconnect "github.com/bautistv/posta-baut/internal/pb/pbconnect"
)

// teamsService implements the Connect TeamsService handler interface
type teamsService struct {
	// Your client instance here
	client client.Client
}

func (s *teamsService) SendMessage(ctx context.Context, req *connect.Request[pb.SendMessageRequest]) (*connect.Response[pb.SendMessageResponse], error) {
	log.Printf("Received message request: %+v", req.Msg)

	// Dummy implementation: just pretend we sent the message
	// Replace with actual client call to send messages
	resp := &pb.SendMessageResponse{
		Success: true,
	}
	return connect.NewResponse(resp), nil
}

func main() {
	// Hardcoded dummy config for client
	cfg := config.ClientConfig{
		MessengerConfig: config.MSGraphClientConfig{
			TenantID: "your-tenant-id",
			ClientID: "your-client-id",
		},
		LookupClientConfig: config.MSGraphClientConfig{
			TenantID: "your-tenant-id",
			ClientID: "your-client-id",
		},
	}

	cli, err := client.NewClient(cfg.MessengerConfig, cfg.LookupClientConfig)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	// Create service handler with client
	teamsSvc := &teamsService{
		client: cli,
	}

	mux := http.NewServeMux()
	mux.Handle(pbconnect.NewTeamsServiceHandler(teamsSvc))

	addr := ":8080"
	log.Printf("Starting example server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
