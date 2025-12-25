package messenger_test

import (
	"context"
	"testing"

	"github.com/bautistv/posta-baut/pkg/messenger"
	"github.com/bautistv/posta-baut/pkg/messenger/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const (
	validChatID    = "chat-123"
	validTeamID    = "team-123"
	validChannelID = "channel-123"
	validContent   = "Hello, World!"
)

var (
	validChannelMessage = messenger.TeamsMessage{
		Type:      messenger.MessageTypeChannel,
		TeamID:    validTeamID,
		ChannelID: validChannelID,
		Content:   validContent,
	}

	validChatMessage = messenger.TeamsMessage{
		Type:    messenger.MessageTypeChat,
		ChatID:  validChatID,
		Content: validContent,
	}
)

func TestSend_Fail(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		m          messenger.Messenger
		msg        messenger.TeamsMessage
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "nil messenger",
			m:    nil,
			msg: messenger.TeamsMessage{
				Content: validContent,
			},
			wantErr:    true,
			wantErrMsg: "messenger cannot be nil",
		},
		// Chat message tests
		{
			name: "missing chat ID",
			m:    mocks.NewMockMessenger(nil),
			msg: messenger.TeamsMessage{
				Type:    messenger.MessageTypeChat,
				Content: "Hello, Chat!",
			},
			wantErr:    true,
			wantErrMsg: "missing chat ID from chat message",
		},
		// Channel message tests
		{
			name: "missing channel ID",
			m:    mocks.NewMockMessenger(nil),
			msg: messenger.TeamsMessage{
				Type:    messenger.MessageTypeChannel,
				TeamID:  validTeamID,
				Content: "Hello, Channel!",
			},
			wantErr:    true,
			wantErrMsg: "missing channel ID from channel message",
		},
		{
			name: "missing team ID",
			m:    mocks.NewMockMessenger(nil),
			msg: messenger.TeamsMessage{
				Type:      messenger.MessageTypeChannel,
				ChannelID: validChannelID,
				Content:   "Hello, Channel!",
			},
			wantErr:    true,
			wantErrMsg: "missing team ID from channel message",
		},
		{
			name: "successful chat message send",
			m:    mocks.NewMockMessenger(nil),
			msg: messenger.TeamsMessage{
				Type:    messenger.MessageTypeChat,
				ChatID:  validChatID,
				Content: "Hello, Chat!",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := messenger.Send(context.Background(), tt.m, tt.msg)
			if tt.wantErrMsg != "" {
				if err == nil {
					t.Errorf("Send() expected error but got nil")
					return
				}

				require.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}

func TestSend_Success(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		msg     messenger.TeamsMessage
		wantErr bool
	}{
		{
			name:    "successful chat message send",
			msg:     validChatMessage,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockMessenger := mocks.NewMockMessenger(ctrl)

			switch tt.msg.Type {
			case messenger.MessageTypeChannel:
				mockMessenger.EXPECT().SendChannelMessage(gomock.Any(), validTeamID, validChannelID, tt.msg).Return(nil)
			case messenger.MessageTypeChat:
				mockMessenger.EXPECT().SendChatMessage(gomock.Any(), validChatID, tt.msg).Return(nil)
			}

			err := messenger.Send(context.Background(), mockMessenger, tt.msg)
			require.NoError(t, err)
		})
	}
}
