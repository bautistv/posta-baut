# Runbook

You can use the example server set up to send a chat message or channel message.

---

## How to run the server

```bash
go run ./cmd/main/main.go
```

By default, it listens on `localhost:8080`.

## Configuration

Create a copy of `config/local.example.yaml` and name it `local.yaml`.
Replace the config values with your real credentials.

### Check server health & readiness

#### How do I check the server health?

The server exposes the standard gRPC Health service.

```bash
grpcurl -plaintext localhost:50051 grpc.health.v1.Health/Check

# {
#   "status": "SERVING"
# }
```

### How do I know the Teams service is available?

```bash
# Inspect the registered services on the server
grpcurl -plaintext localhost:8080 list
# pb.TeamsService

# Inspect the methods belonging to pb.TeamsService
grpcurl -plaintext localhost:8080 list pb.TeamsService
# pb.TeamsService.SendMessage

# Inspect request/response schemas
grpcurl -plaintext localhost:8080 describe pb.TeamsService.SendMessage
# pb.TeamsService.SendMessage is a method:
# rpc SendMessage ( .pb.SendMessageRequest ) returns ( .pb.SendMessageResponse );
```

---

## How to send a message

### Chat Message

```bash
grpcurl -plaintext \
  -d '{
    "target": {
      "chat": {
        "chatId": "your-chat-id"
      }
    },
    "content": "Hello from grpcurl to a Teams chat!"
  }' \
  localhost:50051 \
  pb.TeamsService.SendMessage
```

### Channel Message

```bash
grpcurl -plaintext \
  -d '{
    "target": {
      "channel": {
        "teamId": "your-team-id",
        "channelId": "your-channel-id"
      }
    },
    "content": "Hello from grpcurl to a Teams channel!"
  }' \
  localhost:8080 \
  pb.TeamsService.SendMessage
```
