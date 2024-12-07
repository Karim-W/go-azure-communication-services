package identity

import (
	"context"
	"encoding/json"
)

func (i *_Identity) CreateIdentity(
	ctx context.Context,
	opts *CreateIdentityOptions,
) (*ACSIdentity, error) {
	err := opts.isValid()
	if err != nil {
		return nil, err
	}

	var response createIdentityResponse

	res, err := i.client.Post(
		ctx,
		i.host,
		"/identities",
		map[string][]string{"api-version": {apiVersion}},
		opts,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &ACSIdentity{
		ID:        response.Identity.ID,
		Token:     response.AccessToken.Token,
		ExpiresOn: response.AccessToken.ExpiresOn,
	}, nil
}
