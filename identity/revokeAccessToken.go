package identity

import "context"

func (i *_Identity) RevokeAccessToken(
	ctx context.Context,
	acsId string,
) error {
	_, err := i.client.Post(
		ctx,
		i.host,
		"/identities/"+acsId+"/:revokeAccessTokens",
		map[string][]string{"api-version": {apiVersion}},
		nil,
	)
	return err
}
