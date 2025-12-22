package messenger

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	config "github.com/bautistv/posta-baut/cmd/config"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// GraphMessenger implements the Messenger interface using Microsoft Graph API.
type GraphMessenger struct {
	client msgraphsdk.GraphServiceClient
}

// SendChannelMessage sends a message to a specified channel in a team.
// Returns an error if the operation fails.
func (gm *GraphMessenger) SendChannelMessage(ctx context.Context, teamID string, channelID string, msgContent string) error {
	requestBody := graphmodels.NewChatMessage()
	body := graphmodels.NewItemBody()
	content := msgContent
	body.SetContent(&content)
	requestBody.SetBody(body)

	_, err := gm.client.Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Messages().
		Post(context.Background(), requestBody, nil)

	if err != nil {
		return fmt.Errorf(
			"failed to send channel message to channel id %s belonging to team id %s: %w",
			channelID,
			teamID,
			err,
		)
	}
	return nil
}

// SendChatMessage sends a message to a specified chat.
// Returns an error if the operation fails.
func (gm *GraphMessenger) SendChatMessage(ctx context.Context, chatID string, msgContent string) error {
	requestBody := graphmodels.NewChatMessage()
	body := graphmodels.NewItemBody()
	content := msgContent
	body.SetContent(&content)
	requestBody.SetBody(body)

	_, err := gm.client.Chats().
		ByChatId(chatID).
		Messages().
		Post(context.Background(), requestBody, nil)
	if err != nil {
		return fmt.Errorf("failed to send chat message to chat-id \"%s\": %w", chatID, err)
	}
	return nil
}

func NewGraphMessenger(config config.MSGraphClientConfig) (GraphMessenger, error) {
	// Initialise MS Graph Client
	msGraphClient, err := NewMSGraphClient(config)
	if err != nil {
		return GraphMessenger{}, fmt.Errorf("failed to create MS Graph Messenger: %w", err)
	}

	return GraphMessenger{
		client: msGraphClient,
	}, nil
}

func NewMSGraphClient(config config.MSGraphClientConfig) (msgraphsdk.GraphServiceClient, error) {
	cred, _ := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		TenantID: config.TenantID,
		ClientID: config.ClientID,
		UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
			fmt.Println(message.Message)
			return nil
		},
	})

	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"User.Read"})
	if err != nil {
		fmt.Printf("error creating client: %v\n", err)
		return msgraphsdk.GraphServiceClient{}, fmt.Errorf("Error creating Microsoft Graph Client: %w", err)
	}
	return *client, nil
}
