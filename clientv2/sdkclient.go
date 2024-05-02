package clientv2

import (
	"context"

	"github.com/karim-w/go-azure-communication-services/logger"
)

type client_ struct {
	key    string
	logger logger.Logger
}

func (c *client_) WithLogger(logger logger.Logger) Client {
	c.logger = logger
	return c
}

type Client interface {
	WithLogger(logger logger.Logger) Client
	Get(
		ctx context.Context,
		host string,
		resource string,
		queryParsms map[string][]string,
	) ([]byte, error)
	Delete(
		ctx context.Context,
		host string,
		resource string,
		queryParsms map[string][]string,
	) ([]byte, error)
	Post(
		ctx context.Context,
		host string,
		resource string,
		queryParsms map[string][]string,
		reqbody any,
	) ([]byte, error)
	Put(
		ctx context.Context,
		host string,
		resource string,
		queryParsms map[string][]string,
		reqbody any,
	) ([]byte, error)
	Patch(
		ctx context.Context,
		host string,
		resource string,
		queryParsms map[string][]string,
		reqbody any,
	) ([]byte, error)
}

func New(
	key string,
) *client_ {
	return &client_{key, logger.Noop()}
}

func NewWithLogger(
	key string,
	logger logger.Logger,
) Client {
	return &client_{key, logger}
}
