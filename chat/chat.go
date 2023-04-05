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
}

type _chat struct {
	host       string
	client     *client.Client
	token      string
	validUntil time.Time
	ic         identity.Identity
}

func New(host string, key string) (Chat, error) {
	client := client.New(key)
	ic := identity.New(host, key)
	id, err := ic.CreateIdentity(context.TODO(), &identity.CreateIdentityOptions{
		CreateTokenWithScopes: []string{"chat"},
		ExpiresInMinutes:      60 * 24,
	})
	if err != nil {
		return nil, err
	}
	return &_chat{
		host:       host,
		client:     client,
		token:      id.Token,
		ic:         ic,
		validUntil: time.Now().Add(time.Minute * 60 * 23),
	}, nil
}

func (c *_chat) refreshToken() error {
	id, err := c.ic.CreateIdentity(context.TODO(), &identity.CreateIdentityOptions{
		CreateTokenWithScopes: []string{"chat"},
		ExpiresInMinutes:      60 * 24,
	})
	if err != nil {
		return err
	}
	c.token = id.Token
	c.validUntil = time.Now().Add(time.Minute * 60 * 23)
	return nil
}

func (c *_chat) getToken() (string, error) {
	if time.Now().After(c.validUntil) {
		err := c.refreshToken()
		if err != nil {
			return "", err
		}
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
