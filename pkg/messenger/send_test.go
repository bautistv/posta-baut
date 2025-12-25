package messenger

import (
	"context"
	"testing"

	mocks "github.com/bautistv/posta-baut/pkg/messenger/mocks"
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
		{
			name: "Send Chat Message - Missing Chat ID",
			args: args{
				ctx: context.Background(),
				m:   mocks.NewMockMessenger(nil),
				msg: TeamsMessage{
					Type:    MessageTypeChat,
					Content: "Hello, World!",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Send(tt.args.ctx, tt.args.m, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
