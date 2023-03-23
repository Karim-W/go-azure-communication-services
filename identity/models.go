package identity

import (
	"errors"
	"time"
)

type createIdentityResponse struct {
	Identity struct {
		ID string `json:"id"`
	} `json:"identity"`
	AccessToken struct {
		Token     string    `json:"token"`
		ExpiresOn time.Time `json:"expiresOn"`
	} `json:"accessToken"`
}

type CreateIdentityOptions struct {
	CreateTokenWithScopes []string `json:"createTokenWithScopes"`
	ExpiresInMinutes      int      `json:"expiresInMinutes"`
}

type IssueTokenOptions struct {
	Scopes           []string `json:"scopes"`
	ExpiresInMinutes int      `json:"expiresInMinutes"`
}

type ACSIdentity struct {
	ID        string    `json:"id"`
	Token     string    `json:"token"`
	ExpiresOn time.Time `json:"expiresOn"`
}

const apiVersion = "2022-10-01"

var (
	ERR_NIL_OPTIONS            = errors.New("options cannot be nil")
	ERR_SCOPES_CANNOT_BE_EMPTY = errors.New("scopes cannot be empty")
	ERR_EXPIRY_OUT_OF_RANGE    = errors.New("expiry must be between 60 and 1440 minutes")
)

func (c *CreateIdentityOptions) isValid() error {
	if c == nil {
		return ERR_NIL_OPTIONS
	}
	if c.CreateTokenWithScopes == nil {
		return ERR_SCOPES_CANNOT_BE_EMPTY
	}
	if c.ExpiresInMinutes < 60 || c.ExpiresInMinutes > 1440 {
		return ERR_EXPIRY_OUT_OF_RANGE
	}
	return nil
}

func (i *IssueTokenOptions) isValid() error {
	if i == nil {
		return ERR_NIL_OPTIONS
	}
	if i.Scopes == nil {
		return ERR_SCOPES_CANNOT_BE_EMPTY
	}
	if i.ExpiresInMinutes < 60 || i.ExpiresInMinutes > 1440 {
		return ERR_EXPIRY_OUT_OF_RANGE
	}
	return nil
}

type issueAccessTokenResponse struct {
	Token     string    `json:"token"`
	ExpiresOn time.Time `json:"expiresOn"`
}
