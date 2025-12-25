package msgraph

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func NewMSGraphClient(tenantID string, clientID string) (msgraphsdk.GraphServiceClient, error) {
	if tenantID == "" {
		return msgraphsdk.GraphServiceClient{}, fmt.Errorf("tenant ID is required")
	}
	if clientID == "" {
		return msgraphsdk.GraphServiceClient{}, fmt.Errorf("client ID is required")
	}

	cred, _ := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		TenantID: tenantID,
		ClientID: clientID,
		UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
			fmt.Println(message.Message)
			return nil
		},
	})

	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"User.Read"})
	if err != nil {
		fmt.Printf("error creating client: %v\n", err)
		return msgraphsdk.GraphServiceClient{}, fmt.Errorf("error creating Microsoft Graph Client: %w", err)
	}
	return *client, nil
}
