# go-azure-communication-services

**UNOFFICIAL** GoLang SDK for Azure Communication Services
with support for identity and rooms and ChatThreads.

## usage

get the package

```bash
go get github.com/karim-w/go-azure-communication-services
```

import the package

```go
import (
  "github.com/karim-w/go-azure-communication-services/rooms"
  "github.com/karim-w/go-azure-communication-services/identity"
  "github.com/karim-w/go-azure-communication-services/chat"
)
```

create an identity client

```go

resourceHost := "my-resource.communication.azure.com" // microsoft calls this endpoint
accessKey := "my-access-key"

roomsClient := rooms.NewClient(resourceHost, accessKey)
identityClient := identity.NewClient(resourceHost, accessKey)
```

## identity

### create identity

```go
identity, err := identityClient.CreateIdentity(
context.Background(),
&identity.CreateIdentityOptions{
  CreateTokenWithScopes: []string{"chat"},
  ExpiresInMinutes:      60,
},
)
```

### issue access token

```go
token, err := identityClient.IssueAccessToken(
context.Background(),
acsId,
&identity.IssueAccessTokenOptions{
  Scopes: []string{"chat"},
  ExpiresInMinutes:      60,
},
)
```

### revoke access token

```go
err := identityClient.RevokeAccessToken(
context.Background(),
acsId,
)
```

### delete identity

```go
err := identityClient.DeleteIdentity(
context.Background(),
acsId,
)
```

## rooms

### create room

```go
room, err := roomsClient.CreateRoom(
  context.Background(),
  &rooms.CreateRoomOptions{
    ValidFrom: time.Now(),
    ValidUntil: time.Now().Add(time.Hour * 24),
    RoomJoinPolicy: "InviteOnly",
    Participants: []rooms.RoomParticipant{
      CreateRoomParticipant(id, PRESENTER),
      CreateRoomParticipant(id, ATTENDEE),
    },
  },
)
```

### get room

```go
room, err := roomsClient.GetRoom(
  context.Background(),
  roomId,
)
```

### update room

```go
room, err := roomsClient.UpdateRoom(
  context.Background(),
  roomId,
  &rooms.UpdateRoomOptions{
    ValidFrom: time.Now(),
    ValidUntil: time.Now().Add(time.Hour * 24),
    RoomJoinPolicy: "InviteOnly",
  },
)
```

### delete room

```go
err := roomsClient.DeleteRoom(
  context.Background(),
  roomId,
)
```

### get room participants

```go
participants, err := roomsClient.GetRoomParticipants(
  context.Background(),
  roomId,
)
```

### add room participants

```go
participants, err := roomsClient.AddRoomParticipants(
  context.Background(),
  roomId,
  CreateRoomParticipant(id, PRESENTER),
  CreateRoomParticipant(id, ATTENDEE),
)
```

### remove room participants

```go
err := roomsClient.RemoveRoomParticipants(
  context.Background(),
  roomId,
  RemoveRoomParticipant(id),
)
```

### Update room participants

```go
participants, err := roomsClient.UpdateRoomParticipants(
  context.Background(),
  roomId,
  CreateRoomParticipant(id, PRESENTER),
)
```

> Please Refer to the tests for more examples on how to use the rooms SDK.

## ChatThreads

### Create ChatThread

```go
chatThread, err := chatClient.CreateChatThread(
  context.Background(),
 "test",
 ChatUser{ID: id, DisplayName: "test"},
 ChatUser{ID: id2, DisplayName: "test2"},
)
```

### Delete ChatThread

```go
err := chatClient.DeleteChatThread(
  context.Background(),
  chatThreadId,
)
```

### Add ChatThread Participants

```go
participants, err := chatClient.AddChatParticipants(
  context.Background(),
  chatThreadId,
  ChatUser{ID: id, DisplayName: "test"},
  ChatUser{ID: id2, DisplayName: "test2"},
)
```

### Remove ChatThread Participants

```go
err := chatClient.RemoveChatParticipant(
  context.Background(),
  chatThreadId,
  id,
)
```

## Emails

### Send Email

```go
 client := emails.NewClient(host, key, nil)
 payload := emails.Payload{
  Headers: emails.Headers{
   ClientCorrelationID:    "1234",
   ClientCustomHeaderName: "ClientCustomHeaderValue",
  },
  SenderAddress: "<ACS_EMAIL>"
  Content: emails.Content{
   Subject:   "Test email",
   PlainText: "This is a test email",
  },
  Recipients: emails.Recipients{
   To: []emails.ReplyTo{
    {
     Address: "<EMAIL_ADDRESS>",
    },
   },
  },
 }
 result, err := client.SendEmail(context.TODO(), payload)
```

## References

- [Identity API](https://learn.microsoft.com/en-us/rest/api/communication/communication-identity)
- [Rooms API](https://learn.microsoft.com/en-us/rest/api/communication/rooms)
- [Chat API](https://learn.microsoft.com/en-us/rest/api/communication/chat/chat)
- [Email API](https://learn.microsoft.com/en-us/rest/api/communication/dataplane/email?view=rest-communication-dataplane-2023-03-31)

## License

BSD 3-Clause License

## Author

karim-w

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
