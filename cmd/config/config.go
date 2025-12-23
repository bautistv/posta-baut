package msgraph

// AppConfig holds the overall application configuration.
type AppConfig struct {
	Client ClientConfig `yaml:"client_config"`
}

// ClientConfig holds configuration for a messenger client.
type ClientConfig struct {
	SenderConfig    MSGraphClientConfig `yaml:"sender_config"`
	LookupClientConfig MSGraphClientConfig `yaml:"lookup_client_config"`
}

// MSGraphClientConfig holds credentials for creating a Microsoft Graph Client.
type MSGraphClientConfig struct {
	TenantID string `yaml:"tenant_id"`
	ClientID string `yaml:"client_id"`
}
