# go-azure-communication-services

**UNOFFICIAL** and **NON-STABLE** GoLang SDK for Azure Communication Services
with support for identity and rooms.

> all of the functions **NEED** to be retried in case of failure as testing shows the ACS Azure resource is unstable

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
)
```

create an identity client

```go

resourceHost := "https://my-resource.communication.azure.com" // microsoft calls this endpoint
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
  },
)
```

> Adding participants to the room seem to be broken

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

## References

- [Identity API](https://learn.microsoft.com/en-us/rest/api/communication/communication-identity)
- [Rooms API](https://learn.microsoft.com/en-us/rest/api/communication/rooms)

## License

BSD 3-Clause License

## Author

karim-w

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
