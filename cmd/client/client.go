package client

import (
	"log"

	config "github.com/bautistv/posta-baut/cmd/config"
	lookup "github.com/bautistv/posta-baut/cmd/shared/lookup"
	"github.com/bautistv/posta-baut/pkg/messenger"
)

// NewClient creates a new Client with the provided Messenger and LookupClient configurations.
func NewClient(messengerConfig config.MSGraphClientConfig, lookupClientConfig config.MSGraphClientConfig) (Client, error) {
	messenger, err := messenger.NewGraphMessenger(messengerConfig)
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
