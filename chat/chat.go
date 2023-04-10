package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/karim-w/go-azure-communication-services/client"
	"github.com/karim-w/go-azure-communication-services/identity"
	"github.com/karim-w/stdlib/httpclient"
)

type Chat interface {
	CreateChatThread(
		ctx context.Context,
		topic string,
		participants ...ChatUser,
	) (*CreateChatThreadResponse, error)
	DeleteChatThread(
		ctx context.Context,
		threadID string,
	) error
	WithToken(
		token string,
		ExpiresAt time.Time,
	) Chat
}

type _chat struct {
	host       string
	client     *client.Client
	token      string
	validUntil time.Time
}

func New(host string, key string) (Chat, error) {
	client := client.New(key)
	return &_chat{
		host:   host,
		client: client,
		token:  "",
	}, nil
}

func NewWithToken(host string, token string, expiresAt time.Time) (Chat, error) {
	return &_chat{
		host:       host,
		client:     nil,
		token:      token,
		validUntil: expiresAt,
	}, nil
}

func (c *_chat) getToken() (string, error) {
	if time.Now().After(c.validUntil) {
		return "", ERR_EXPIRED_TOKEN
	}
	if c.token == "" {
		return "", ERR_NO_TOKEN_PROVIDED
	}
	return c.token, nil
}

func (c *_chat) CreateChatThread(
	ctx context.Context,
	topic string,
	participants ...ChatUser,
) (*CreateChatThreadResponse, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	req := CreateChatThread{
		Topic: topic,
	}
	for _, p := range participants {
		req.Participants = append(req.Participants, Participant{
			CommunicationIdentifier: identity.CommunicationIdentifier{
				RawID: p.ID,
				CommunicationUser: identity.CommunicationUser{
					ID: p.ID,
				},
			},
			DisplayName: p.DisplayName,
		})
	}
	response := CreateChatThreadResponse{}
	res := httpclient.Req(
		"https://"+c.host+"/chat/threads?api-version="+_apiVersion,
	).AddBearerAuth(
		token,
	).AddHeader("Content-Type", "application/json").
		AddBody(req).Post()
	if res.IsSuccess() {
		err := res.SetResult(&response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	}
	err = fmt.Errorf(string(res.GetBody()))
	return nil, err
}

func (c *_chat) DeleteChatThread(
	ctx context.Context,
	threadID string,
) error {
	token, err := c.getToken()
	if err != nil {
		return err
	}

	res := httpclient.Req(
		"https://"+c.host+"/chat/threads/"+threadID+"?api-version="+_apiVersion,
	).AddBearerAuth(
		token,
	).AddHeader("Content-Type", "application/json").
		Del()
	if res.IsSuccess() {
		return nil
	}
	err = fmt.Errorf(string(res.GetBody()))
	return err
}

func (c *_chat) WithToken(
	token string,
	ExpiresAt time.Time,
) Chat {
	c.token = token
	c.validUntil = ExpiresAt
	return c
}
