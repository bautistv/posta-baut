package client

import (
	"reflect"
	"testing"

	config "github.com/bautistv/posta-baut/cmd/config"
)

func TestNewClient(t *testing.T) {
	type args struct {
		messengerConfig    config.MSGraphClientConfig
		lookupClientConfig config.MSGraphClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.messengerConfig, tt.args.lookupClientConfig)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
