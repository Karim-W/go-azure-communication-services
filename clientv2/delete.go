package clientv2

import (
	"context"
	"errors"

	"github.com/karim-w/go-azure-communication-services/logger"
)

func (c *client_) Delete(
	ctx context.Context,
	host string,
	resource string,
	queryParsms map[string][]string,
) ([]byte, error) {
	req, err := c.buildRequest("DELETE", host, resource, queryParsms, nil)
	if err != nil {
		c.logger.Error("DELETE Request failed", logger.String("error", err.Error()))
		return nil, err
	}

	res := req.Del()

	responseBody := res.GetBody()

	if !res.IsSuccess() {
		code := res.GetStatusCode()

		c.logger.Error(
			"DELETE Request failed",
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
