package msgraph

// AppConfig holds configuration for a messenger client.
type AppConfig struct {
	SenderConfig       ClientConfig `yaml:"sender_config"`
	LookupClientConfig ClientConfig `yaml:"lookup_client_config"`
}

// ClientConfig holds credentials for creating a Microsoft Graph Client.
type ClientConfig struct {
	TenantID string `yaml:"tenant_id"`
	ClientID string `yaml:"client_id"`
}
