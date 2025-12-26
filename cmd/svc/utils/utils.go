// Package utils provides service-level utilities for the Posta Baut application.
package utils

import (
	"fmt"

	pbv1 "github.com/bautistv/posta-baut/internal/pb/v1"
	"github.com/bautistv/posta-baut/pkg/messenger"

	"github.com/google/uuid"
	"github.com/zeebo/xxh3"
)

// ReqToMsg converts the protobuf SendMessageRequest into the domain SendMessageRequest.
func ReqToMsg(pbReq *pbv1.SendMessageRequest) (*messenger.SendMessageRequest, error) {
    if pbReq == nil {
        return nil, fmt.Errorf("request cannot be nil")
    }

    if pbReq.Target == nil {
        return nil, fmt.Errorf("target is required")
    }

    var target messenger.MessageTarget

    switch t := pbReq.Target.Target.(type) {
    case *pbv1.MessageTarget_Channel:
        if t.Channel == nil {
            return nil, fmt.Errorf("channel target is nil")
        }
        target.Channel = &messenger.TeamsChannelTarget{
            TeamID:    t.Channel.GetTeamId(),
            ChannelID: t.Channel.GetChannelId(),
            ThreadID:  t.Channel.GetThreadId(),
        }
    case *pbv1.MessageTarget_Chat:
        if t.Chat == nil {
            return nil, fmt.Errorf("chat target is nil")
        }
        target.Chat = &messenger.ChatTarget{
            ChatID:          t.Chat.GetChatId(),
            ReplyToMessageID: t.Chat.GetReplyToMessageId(),
        }
    default:
        return nil, fmt.Errorf("unknown message target type: %T", t)
    }

    if pbReq.Content == "" {
        return nil, fmt.Errorf("content cannot be empty")
    }

    domainReq := &messenger.SendMessageRequest{
        Target:  target,
        Content: pbReq.GetContent(),
    }

    return domainReq, nil
}

func MsgToUID(msg messenger.SendMessageRequest) uuid.UUID {
  var hash [16]byte
  var guid uuid.UUID

  input := fmt.Sprintf("%s%s%s%s%s", msg.Target.Channel.ChannelID,
  	msg.Target.Channel.TeamID,
	msg.Target.Chat.ChatID,
	msg.Target.Channel.ThreadID,
	msg.Target.Chat.ReplyToMessageID,
	)
  hash = xxh3.HashString128(input).Bytes()

  // uuid.FromBytes returns an error if the slice
  // of bytes is not 16 - as hash is defined as
  // [16]byte then we can ignore checking the error
  guid, _ = uuid.FromBytes(hash[:])
  return guid
}
