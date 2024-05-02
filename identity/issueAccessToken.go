package identity

import (
	"context"
	"encoding/json"
)

func (i *_Identity) IssueAccessToken(
	ctx context.Context,
	acsId string,
	opts *IssueTokenOptions,
) (*ACSIdentity, error) {
	if err := opts.isValid(); err != nil {
		return nil, err
	}

	var response issueAccessTokenResponse

	res, err := i.client.Post(
		ctx,
		i.host,
		"/identities/"+acsId+"/:issueAccessToken",
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
		ID:        acsId,
		Token:     response.Token,
		ExpiresOn: response.ExpiresOn,
	}, nil
}
