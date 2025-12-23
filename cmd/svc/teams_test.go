// Package svc provides service-level utilities for the Posta Baut application.
package svc

import (
	"context"
	"reflect"
	"testing"

	"connectrpc.com/connect"
	"github.com/bautistv/posta-baut/cmd/client"
	pb "github.com/bautistv/posta-baut/internal/pb/v1"
)

func TestNewTeamsServiceClient(t *testing.T) {
	type args struct {
		client *client.Client
	}
	tests := []struct {
		name string
		args args
		want *teamsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTeamsServiceClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTeamsServiceClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamsService_SendMessage(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[pb.SendMessageRequest]
	}
	tests := []struct {
		name    string
		s       *teamsService
		args    args
		want    *connect.Response[pb.SendMessageResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SendMessage(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("teamsService.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamsService.SendMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
