// Package utils provides service-level utilities for the Posta Baut application.
package utils

import (
	"reflect"
	"testing"

	pb "github.com/bautistv/posta-baut/internal/pb/v1"
	"github.com/bautistv/posta-baut/pkg/messenger"
)

func TestReqToMsg(t *testing.T) {
	type args struct {
		req *pb.SendMessageRequest
	}
	tests := []struct {
		name    string
		args    args
		want    messenger.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReqToMsg(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ReqToMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqToMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
