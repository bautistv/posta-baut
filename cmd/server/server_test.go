package server

import (
	"testing"

	"github.com/bautistv/posta-baut/cmd/client"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	t.Run("successful server initialisation", func(t *testing.T) {
		teamsServiceClient := client.Client{}
		_, err := NewServer(teamsServiceClient)
		require.NoErrorf(t, err, "NewServer() error = %v, expected no error", err)
	})
}

// TODO: TestServerEndpoints tests that all expected endpoints are registered - https://github.com/bautistv/posta-baut/issues/12
func TestServerEndpoints(t *testing.T) {
}

// TODO: TestServerGracefulShutdown tests the graceful shutdown mechanism - https://github.com/bautistv/posta-baut/issues/12
func TestServerGracefulShutdown(t *testing.T) {
}

// TODO: TestServerShutdownTimeout tests shutdown timeout behavior - https://github.com/bautistv/posta-baut/issues/12
func TestServerShutdownTimeout(t *testing.T) {
}

// TODO: TestServerConfiguration tests server configuration parameters - https://github.com/bautistv/posta-baut/issues/12
func TestServerConfiguration(t *testing.T) {
}
