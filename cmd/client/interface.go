package client

import (
	"github.com/bautistv/posta-baut/cmd/messenger"
	lookup "github.com/bautistv/posta-baut/cmd/shared/lookup"
)

type Client struct {
	Messenger messenger.Messenger
	LookupClient lookup.TeamsLookupClient
}
