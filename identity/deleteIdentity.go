package identity

import "context"

func (i *_Identity) DeleteIdentity(
	ctx context.Context,
	acsId string,
) error {
	_, err := i.client.Delete(
		ctx,
		i.host,
		"/identities/"+acsId,
		map[string][]string{"api-version": {apiVersion}},
	)
	return err
}
