package client

import (
	"github.com/bautistv/posta-baut/cmd/messenger"
	lookup "github.com/bautistv/posta-baut/cmd/shared/lookup"
)

func NewClient(messenger messenger.Messenger, lookupClient lookup.TeamsLookupClient) *Client {
	return &Client{
		Messenger:    messenger,
		LookupClient: lookupClient,
	}
}
