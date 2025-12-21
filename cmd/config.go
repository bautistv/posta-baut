package posta_baut

type ClientConfig struct {
	msClientConfig MSGraphClientConfig
	// Add configuration fields as needed
}

// MSGraphClientConfig holds credentials for creating a Microsoft Graph Client.
// See https://learn.microsoft.com/en-us/graph/sdks/create-client?from=snippets&tabs=go
type MSGraphClientConfig struct {
	TenantID     string
	ClientID     string
	ClientSecret string
}
