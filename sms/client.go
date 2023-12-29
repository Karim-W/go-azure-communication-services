package sms

import (
	"context"

	"github.com/karim-w/go-azure-communication-services/client"
)

type Client interface {
	SendSMS(
		ctx context.Context,
		version *string,
		payload Request,
	) (Result, error)
}

type _client struct {
	cl   *client.Client
	host string
}

func NewClient(
	host string,
	key string,
	version *string,
) Client {
	return &_client{
		cl:   client.New(key),
		host: host,
	}
}

// SendSMS() sends a SMS using the Azure Communication Services REST API.
// Parameters:
//
//   - ctx: The context of the request.
//   - version: The version of the API to use. If nil, the default API version will be used.
//   - payload: The payload of the request check the models.go file for more information.
//
// example:
//
//	result, err := client.SendSMS(context.TODO(),nil, Request{})
func (c *_client) SendSMS(
	ctx context.Context,
	version *string,
	payload Request,
) (Result, error) {
	var result Result

	v := "api-version="
	if version == nil {
		v += defaultAPIVersion
	} else {
		v += *version
	}

	err := c.cl.Post(ctx, c.host, "/sms", v, payload, &result)

	return result, err
}
