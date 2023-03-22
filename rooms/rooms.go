package rooms

import "context"

type Rooms interface {
	CreateRoom(
		ctx context.Context,
		options *CreateRoomOptions,
	) (*RoomModel, error)
}

type _RoomsClient struct {
	url string
}

func New(url string) Rooms {
	return &_RoomsClient{url}
}

func (c *_RoomsClient) CreateRoom(
	ctx context.Context,
	options *CreateRoomOptions,
) (*RoomModel, error) {
	// ...
	return nil, nil
}
