package messenger

import (
	"context"
	"testing"
)

func TestSend(t *testing.T) {
	type args struct {
		ctx context.Context
		m   Messenger
		msg TeamsMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Send(tt.args.ctx, tt.args.m, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
