package msgraph

import (
	"reflect"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func TestNewMSGraphClient(t *testing.T) {
	type args struct {
		tenantID string
		clientID string
		clientSecret string
	}
	tests := []struct {
		name    string
		args    args
		want    msgraphsdk.GraphServiceClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMSGraphClient(tt.args.tenantID, tt.args.clientID, tt.args.clientSecret)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewMSGraphClient() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMSGraphClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
