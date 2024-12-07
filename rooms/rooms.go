package rooms

import (
	"context"
	"encoding/json"

	"github.com/karim-w/go-azure-communication-services/clientv2"
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
	client clientv2.Client
}

func New(
	host string,
	key string,
) Rooms {
	client := clientv2.New(key)
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
	response, err := c.client.Post(
		ctx,
		c.host,
		"/rooms",
		map[string][]string{
			"api-version": {apiVersion},
		},
		options,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &responseModel)
	if err != nil {
		return nil, err
	}

	return responseModel, nil
}

func (c *_RoomsClient) GetRoom(
	ctx context.Context,
	roomId string,
) (result *RoomModel, err error) {
	res, err := c.client.Get(
		ctx,
		c.host,
		"/rooms/"+roomId,
		map[string][]string{
			"api-version": {apiVersion},
		},
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return
}

func (c *_RoomsClient) UpdateRoom(
	ctx context.Context,
	roomId string,
	options *UpdateRoomOptions,
) (res *RoomModel, err error) {
	if options == nil {
		return nil, ERR_ROOMS_NIL_OPTIONS
	}

	body, err := c.client.Patch(
		ctx,
		c.host,
		"/rooms/"+roomId,
		map[string][]string{
			"api-version": {apiVersion},
		},
		RoomModel{
			ValidFrom:      options.ValidFrom,
			ValidUntil:     options.ValidUntil,
			Participants:   options.Participants,
			RoomJoinPolicy: options.RoomJoinPolicy,
		},
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return
}

func (c *_RoomsClient) DeleteRoom(
	ctx context.Context,
	roomId string,
) error {
	_, err := c.client.Delete(
		ctx,
		c.host,
		"/rooms/"+roomId,
		map[string][]string{
			"api-version": {apiVersion},
		},
	)

	return err
}

func (c *_RoomsClient) AddParticipants(
	ctx context.Context,
	roomId string,
	Participants ...RoomParticipant,
) (*[]RoomParticipant, error) {
	responseModel := &roomParticipantsUpdate{}

	body, err := c.client.Post(
		ctx,
		c.host,
		"/rooms/"+roomId+"/participants:add",
		map[string][]string{
			"api-version": {apiVersion},
		},
		roomParticipantsUpdate{Participants},
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel.Participants, nil
}

func (c *_RoomsClient) GetParticipants(
	ctx context.Context,
	roomId string,
) (*[]RoomParticipant, error) {
	responseModel := &roomParticipantsUpdate{}
	body, err := c.client.Get(
		ctx,
		c.host,
		"/rooms/"+roomId+"/participants",
		map[string][]string{
			"api-version": {apiVersion},
		},
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel.Participants, nil
}

func (c *_RoomsClient) UpdateParticipants(
	ctx context.Context,
	roomId string,
	Participants ...RoomParticipant,
) (*[]RoomParticipant, error) {
	responseModel := &roomParticipantsUpdate{}
	body, err := c.client.Post(
		ctx,
		c.host,
		"/rooms/"+roomId+"/participants:update",
		map[string][]string{
			"api-version": {apiVersion},
		},

		roomParticipantsUpdate{Participants},
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel.Participants, nil
}

func (c *_RoomsClient) RemoveParticipants(
	ctx context.Context,
	roomId string,
	Participants ...RoomParticipant,
) (*[]RoomParticipant, error) {
	responseModel := &roomParticipantsUpdate{}
	body, err := c.client.Post(
		ctx,
		c.host,
		"/rooms/"+roomId+"/participants:remove",
		map[string][]string{
			"api-version": {apiVersion},
		},
		roomParticipantsUpdate{Participants},
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel.Participants, nil
}
