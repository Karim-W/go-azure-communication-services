package rooms

import (
	"context"

	"github.com/karim-w/stdlib/httpclient"
)

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
	if options == nil {
		return nil, ERR_ROOMS_CREATE_ROOM_NIL_OPTIONS
	}
	res := httpclient.Req(c.url+"/rooms?api-version="+apiVersion).
		AddBody(options).
		AddHeader("Content-Type", "application/json").
		Post()
	if !res.IsSuccess() {
		return nil, res.CatchError()
	}
	responseModel := &RoomModel{}
	err := res.SetResult(responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}
