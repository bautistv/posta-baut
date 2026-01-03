package graphhelper

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	auth "github.com/microsoftgraph/msgraph-sdk-go-core/authentication"
)

type GraphHelper struct {
	clientSecretCredential *azidentity.ClientSecretCredential
	appClient              *msgraphsdk.GraphServiceClient
}

func NewGraphHelper() *GraphHelper {
	g := &GraphHelper{}
	return g
}

// NewGraphForAppAuth requests an access token by using the client credentials flow.
// Returns a Graph client using the request adapter.
// Adapted from example method InitializeGraphForAppAuth as of 3 Jan 2026, shown in
// https://learn.microsoft.com/en-us/graph/tutorials/go-app-only-authentication
func (g *GraphHelper) NewGraphForAppAuth(clientId, tenantId, clientSecret string) (*msgraphsdk.GraphServiceClient, error) {
	credential, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
	if err != nil {
		return nil, err
	}

	g.clientSecretCredential = credential

	// Create an auth provider using the credential
	scopes := []string{"https://graph.microsoft.com/.default"}
	authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(g.clientSecretCredential, []string{
		"https://graph.microsoft.com/.default",
	})
	if err != nil {
		return nil, fmt.Errorf("error creating auth provider with required scopes %s: %w", scopes, err)
	}

	// Create a request adapter using the auth provider
	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
		return nil, fmt.Errorf("error creating graph request adapter with auth provider %s: %w", *authProvider, err)
	}

	// Create a Graph client using request adapter
	client := msgraphsdk.NewGraphServiceClient(adapter)

	return client, nil
}
