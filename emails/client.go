package emails

import (
	"context"

	"github.com/karim-w/go-azure-communication-services/client"
)

type Client interface {
	SendEmail(
		ctx context.Context,
		payload Payload,
	) (EmailResult, error)
}

type _client struct {
	cl      *client.Client
	host    string
	version string
}

func NewClient(
	host string,
	key string,
	version *string,
) Client {
	cl := client.New(key)
	v := defaultAPIVersion
	if version != nil {
		v = *version
	}
	apiVersion := "api-version=" + v
	return &_client{
		cl:      cl,
		host:    host,
		version: apiVersion,
	}
}

// SendEmail sends an email using the Azure Communication Services REST API.
// Parameters:
//   - ctx: The context of the request.
//   - payload: The payload of the request check the models.go file for more information.
func (c *_client) SendEmail(
	ctx context.Context,
	payload Payload,
) (EmailResult, error) {
	var result EmailResult
	err := c.cl.Post(ctx, c.host, "/emails:send", c.version, payload, &result)
	return result, err
}
