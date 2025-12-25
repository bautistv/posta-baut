package client

import (
	"log"

	config "github.com/bautistv/posta-baut/cmd/config"
	lookup "github.com/bautistv/posta-baut/cmd/shared/lookup"
	msgraph "github.com/bautistv/posta-baut/pkg/messenger/msgraph"
)

// NewClient creates a new Client with the provided Messenger and LookupClient configurations.
func NewClient(messengerConfig config.ClientConfig, lookupClientConfig config.ClientConfig) (Client, error) {
	messenger, err := msgraph.NewGraphMessenger(messengerConfig)
	if err != nil {
		log.Printf("failed to create Graph Messenger: %v", err)
		return Client{}, err
	}

	lookupClient, err := lookup.NewMSGraphLookupClient(lookupClientConfig)
	if err != nil {
		log.Printf("failed to create MS Graph Lookup Client: %v", err)
		return Client{}, err
	}

	return Client{
		Messenger:    messenger,
		LookupClient: lookupClient,
	}, nil
}
