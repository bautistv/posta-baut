package msgraph

// ClientConfig holds configuration for a messenger client.
type ClientConfig struct {
	MessengerConfig MSGraphClientConfig `yaml:"messenger_config"`
	LookupClientConfig  MSGraphClientConfig `yaml:"lookup_client_config"`
}

// MSGraphClientConfig holds credentials for creating a Microsoft Graph Client.
type MSGraphClientConfig struct {
	TenantID string `yaml:"tenant_id"`
	ClientID string `yaml:"client_id"`
}
