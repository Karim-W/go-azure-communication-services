package identity

import (
	"context"

	"github.com/karim-w/go-azure-communication-services/client"
)

type Identity interface {
	CreateIdentity(
		ctx context.Context,
		opts *CreateIdentityOptions,
	) (*ACSIdentity, error)
	IssueAccessToken(
		ctx context.Context,
		acsId string,
		opts *IssueTokenOptions,
	) (*ACSIdentity, error)
	DeleteIdentity(
		ctx context.Context,
		acsId string,
	) error
	RevokeAccessToken(
		ctx context.Context,
		acsId string,
	) error
}

type _Identity struct {
	client *client.Client
	host   string
}

func New(
	host string,
	key string,
) Identity {
	client := client.New(key)
	return &_Identity{
		client: client,
		host:   host,
	}
}

func (i *_Identity) CreateIdentity(
	ctx context.Context,
	opts *CreateIdentityOptions,
) (*ACSIdentity, error) {
	err := opts.isValid()
	if err != nil {
		return nil, err
	}
	var response createIdentityResponse
	err = i.client.Post(
		ctx,
		i.host,
		"/identities",
		"api-version="+apiVersion,
		opts,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &ACSIdentity{
		ID:        response.Identity.ID,
		Token:     response.AccessToken.Token,
		ExpiresOn: response.AccessToken.ExpiresOn,
	}, nil
}

func (i *_Identity) IssueAccessToken(
	ctx context.Context,
	acsId string,
	opts *IssueTokenOptions,
) (*ACSIdentity, error) {
	if err := opts.isValid(); err != nil {
		return nil, err
	}
	var response issueAccessTokenResponse
	err := i.client.Post(
		ctx,
		i.host,
		"/identities/"+acsId+"/:issueAccessToken",
		"api-version="+apiVersion,
		opts,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &ACSIdentity{
		ID:        acsId,
		Token:     response.Token,
		ExpiresOn: response.ExpiresOn,
	}, nil
}

func (i *_Identity) RevokeAccessToken(
	ctx context.Context,
	acsId string,
) error {
	return i.client.Post(
		ctx,
		i.host,
		"/identities/"+acsId+"/:revokeAccessTokens",
		"api-version="+apiVersion,
		nil,
		nil,
	)
}

func (i *_Identity) DeleteIdentity(
	ctx context.Context,
	acsId string,
) error {
	return i.client.Delete(
		ctx,
		i.host,
		"/identities/"+acsId,
		"api-version="+apiVersion,
		nil,
	)
	// return nil
}
