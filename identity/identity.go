package identity

import (
	"context"

	"github.com/karim-w/go-azure-communication-services/clientv2"
	"github.com/karim-w/go-azure-communication-services/logger"
)

type Identity interface {
	CreateIdentity(
		ctx context.Context,
		opts *CreateIdentityOptions,
	) (*ACSIdentity, error)
	IssueAccessToken(
		ctx context.Context,
		acsId string,
		opts *IssueTokenOptions,
	) (*ACSIdentity, error)
	DeleteIdentity(
		ctx context.Context,
		acsId string,
	) error
	RevokeAccessToken(
		ctx context.Context,
		acsId string,
	) error
}

type _Identity struct {
	client clientv2.Client
	host   string
}

func New(
	host string,
	key string,
) Identity {
	client := clientv2.New(key)
	return &_Identity{
		client: client,
		host:   host,
	}
}

func NewWithLogger(
	host string,
	key string,
	Logger logger.Logger,
) Identity {
	if Logger == nil {
		Logger = logger.Default()
	}

	client := clientv2.NewWithLogger(key, Logger)

	return &_Identity{
		client: client,
		host:   host,
	}
}
