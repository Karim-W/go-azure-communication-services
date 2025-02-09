package emails

import (
	"context"
	"encoding/json"

	"github.com/karim-w/go-azure-communication-services/clientv2"
	"github.com/karim-w/go-azure-communication-services/logger"
)

type Client interface {
	SendEmail(
		ctx context.Context,
		payload Payload,
	) (EmailResult, error)
}

type _client struct {
	cl      clientv2.Client
	host    string
	version string
}

func NewClient(
	host string,
	key string,
	version *string,
) Client {
	cl := clientv2.New(key)
	
	v := defaultAPIVersion
	if version != nil {
		v = *version
	}
	return &_client{
		cl:      cl,
		host:    host,
		version: v,
	}
}

func NewClientWithLogger(
	host string,
	key string,
	version *string,
	Logger logger.Logger,
	
) Client {

	if Logger == nil {
		Logger = logger.Default()
	}
	cl := clientv2.NewWithLogger(key , Logger)	

	v := defaultAPIVersion
	if version != nil {
		v = *version
	}
	return &_client{
		cl:      cl,
		host:    host,
		version: v,
	}
}

// SendEmail sends an email using the Azure Communication Services REST API.
// Parameters:
//   - ctx: The context of the request.
//   - payload: The payload of the request check the models.go file for more information.
func (c *_client) SendEmail(
	ctx context.Context,
	payload Payload,
) (result EmailResult, err error) {
	res, err := c.cl.Post(
		ctx,
		c.host,
		"/emails:send",
		map[string][]string{
			"api-version": {c.version},
		},
		payload,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return
}
