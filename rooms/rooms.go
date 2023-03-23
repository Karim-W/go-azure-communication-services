package rooms

import (
	"context"

	"github.com/karim-w/go-azure-communication-services/client"
)

type Rooms interface {
	CreateRoom(
		ctx context.Context,
		options *CreateRoomOptions,
	) (*RoomModel, error)
	GetRoom(
		ctx context.Context,
		roomId string,
	) (*RoomModel, error)
	UpdateRoom(
		ctx context.Context,
		roomId string,
		options *UpdateRoomOptions,
	) (*RoomModel, error)
	DeleteRoom(
		ctx context.Context,
		roomId string,
	) error
	AddParticipants(
		ctx context.Context,
		roomId string,
		Participants ...RoomParticipant,
	) (*[]RoomParticipant, error)
	GetParticipants(
		ctx context.Context,
		roomId string,
	) (*[]RoomParticipant, error)
	UpdateParticipants(
		ctx context.Context,
		roomId string,
		Participants ...RoomParticipant,
	) (*[]RoomParticipant, error)
	RemoveParticipants(
		ctx context.Context,
		roomId string,
		Participants ...RoomParticipant,
	) (*[]RoomParticipant, error)
}

type _RoomsClient struct {
	host   string
	client *client.Client
}

func New(
	host string,
	key string,
) Rooms {
	client := client.New(key)
	return &_RoomsClient{host, client}
}

func (c *_RoomsClient) CreateRoom(
	ctx context.Context,
	options *CreateRoomOptions,
) (*RoomModel, error) {
	if options == nil {
		return nil, ERR_ROOMS_NIL_OPTIONS
	}
	responseModel := &RoomModel{}
	err := c.client.Post(
		ctx,
		c.host,
		"/rooms",
		"api-version="+apiVersion,
		options,
		&responseModel,
	)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (c *_RoomsClient) GetRoom(
	ctx context.Context,
	roomId string,
) (*RoomModel, error) {
	responseModel := &RoomModel{}
	err := c.client.Get(
		ctx,
		c.host,
		"/rooms/"+roomId,
		"api-version="+apiVersion,
		&responseModel,
	)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (c *_RoomsClient) UpdateRoom(
	ctx context.Context,
	roomId string,
	options *UpdateRoomOptions,
) (*RoomModel, error) {
	if options == nil {
		return nil, ERR_ROOMS_NIL_OPTIONS
	}
	responseModel := &RoomModel{}
	err := c.client.Patch(
		ctx,
		c.host,
		"/rooms/"+roomId,
		"api-version="+apiVersion,
		options,
		&responseModel,
	)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (c *_RoomsClient) DeleteRoom(
	ctx context.Context,
	roomId string,
) error {
	return c.client.Delete(
		ctx,
		c.host,
		"/rooms/"+roomId,
		"api-version="+apiVersion,
		nil,
	)
}

func (c *_RoomsClient) AddParticipants(
	ctx context.Context,
	roomId string,
	Participants ...RoomParticipant,
) (*[]RoomParticipant, error) {
	responseModel := &[]RoomParticipant{}
	err := c.client.Post(
		ctx,
		c.host,
		"/rooms/"+roomId+"/participants:add",
		"api-version="+apiVersion,
		Participants,
		&responseModel,
	)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (c *_RoomsClient) GetParticipants(
	ctx context.Context,
	roomId string,
) (*[]RoomParticipant, error) {
	responseModel := &[]RoomParticipant{}
	err := c.client.Get(
		ctx,
		c.host,
		"/rooms/"+roomId+"/participants",
		"api-version="+apiVersion,
		&responseModel,
	)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (c *_RoomsClient) UpdateParticipants(
	ctx context.Context,
	roomId string,
	Participants ...RoomParticipant,
) (*[]RoomParticipant, error) {
	responseModel := &[]RoomParticipant{}
	err := c.client.Post(
		ctx,
		c.host,
		"/rooms/"+roomId+"/participants:update",
		"api-version="+apiVersion,
		Participants,
		&responseModel,
	)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (c *_RoomsClient) RemoveParticipants(
	ctx context.Context,
	roomId string,
	Participants ...RoomParticipant,
) (*[]RoomParticipant, error) {
	responseModel := &[]RoomParticipant{}
	err := c.client.Post(
		ctx,
		c.host,
		"/rooms/"+roomId+"/participants:remove",
		"api-version="+apiVersion,
		Participants,
		&responseModel,
	)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}
