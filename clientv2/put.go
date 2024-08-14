package clientv2

import (
	"context"
	"errors"

	"github.com/karim-w/go-azure-communication-services/logger"
)

func (c *client_) Put(
	ctx context.Context,
	host string,
	resource string,
	queryParsms map[string][]string,
	reqbody any,
) ([]byte, error) {
	req, err := c.buildRequest("PUT", host, resource, queryParsms, reqbody)
	if err != nil {
		c.logger.Error("PUT Request failed", logger.String("error", err.Error()))
		return nil, err
	}

	res := req.Put()

	responseBody := res.GetBody()

	if !res.IsSuccess() {
		code := res.GetStatusCode()

		c.logger.Error(
			"PUT Request failed",
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
