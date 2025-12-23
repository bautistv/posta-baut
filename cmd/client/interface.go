package client

import (
	lookup "github.com/bautistv/posta-baut/cmd/shared/lookup"
	"github.com/bautistv/posta-baut/pkg/messenger"
)

type Client struct {
	Messenger    messenger.Messenger
	LookupClient lookup.TeamsLookupClient
}
