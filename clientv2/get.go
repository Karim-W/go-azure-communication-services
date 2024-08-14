package clientv2

import (
	"context"
	"errors"

	"github.com/karim-w/go-azure-communication-services/logger"
)

func (c *client_) Get(
	ctx context.Context,
	host string,
	resource string,
	queryParsms map[string][]string,
) ([]byte, error) {
	req, err := c.buildRequest("GET", host, resource, queryParsms, nil)
	if err != nil {
		c.logger.Error("GET Request failed", logger.String("error", err.Error()))
		return nil, err
	}

	res := req.Get()

	responseBody := res.GetBody()

	if !res.IsSuccess() {
		code := res.GetStatusCode()

		c.logger.Error(
			"GET Request failed",
			logger.String("response", string(responseBody)),
			logger.Int("status_code", code),
			logger.String("CURL", res.CURL()),
		)

		return nil, errors.New(string(responseBody))
	}

	if len(responseBody) == 0 {
		return nil, nil
	}

	return responseBody, nil
}
