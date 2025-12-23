package client

import (
	lookup "github.com/bautistv/posta-baut/cmd/shared/lookup"
	"github.com/bautistv/posta-baut/pkg/messenger"
)

// Client encapsulates the Messenger and LookupClient for interacting with Microsoft Teams.
type Client struct {
	Messenger    messenger.Messenger
	LookupClient lookup.TeamsLookupClient
}
