package calls

import (
	"context"
	"fmt"
	"time"

	"github.com/karim-w/go-azure-communication-services/identity"
)

func (c *_calls) refreshToken() error {
	client := *c.idc
	user, err := client.IssueAccessToken(
		context.Background(),
		c.id,
		&identity.IssueTokenOptions{
			ExpiresInMinutes: 1440,
			Scopes:           []string{"chat", "void"},
		},
	)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("failed to create identity")
	}
	c.token = user.Token
	c.validUntil = user.ExpiresOn
	return nil
}

func NewWithToken(host string, token string, expiresAt time.Time) (Call, error) {
	return &_calls{
		host:       host,
		client:     nil,
		token:      token,
		validUntil: expiresAt,
	}, nil
}

func (c *_calls) getToken() (string, error) {
	if time.Now().After(c.validUntil) {
		if c.idc != nil && c.id != "" {
			err := c.refreshToken()
			if err != nil {
				return c.token, err
			}
			return "", ERR_EXPIRED_TOKEN
		}
	}

	if c.token == "" {
		return "", ERR_NO_TOKEN_PROVIDED
	}

	return c.token, nil
}

func (c *_calls) GetToken() (string, error) {
	if c.tokenFetcher != nil {
		return (*c.tokenFetcher)()
	}

	return c.getToken()
}

func (c *_calls) SetTokenFetcher(
	fetcher func() (string, error),
) {
	c.tokenFetcher = &fetcher
}

func (c *_calls) WithToken(
	token string,
	ExpiresAt time.Time,
) Call {
	c.token = token
	c.validUntil = ExpiresAt
	return c
}
