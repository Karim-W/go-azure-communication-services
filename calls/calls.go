package calls

import (
	"context"
	"fmt"
	"time"

	"github.com/karim-w/go-azure-communication-services/clientv2"
	"github.com/karim-w/go-azure-communication-services/identity"
	"github.com/karim-w/go-azure-communication-services/logger"
)

type Call interface {
	CreateOutBoundCall(
		ctx context.Context,
		options *CreateOutboundCallOptions,
	) (*OutboundCall, error)
	AddParticipant(
		ctx context.Context,
		callId string,
		info AddPraticipantRequest,
	) error
	WithToken(
		token string,
		ExpiresAt time.Time,
	) Call
	GetToken() (string, error)
	SetTokenFetcher(
		fetcher func() (string, error),
	)
}

type _calls struct {
	host         string
	client       clientv2.Client
	token        string
	validUntil   time.Time
	idc          *identity.Identity
	id           string
	tokenFetcher *func() (string, error)
}

func New(host string, key string) (Call, error) {
	identityClient := identity.New(host, key)
	user, err := identityClient.CreateIdentity(
		context.Background(),
		&identity.CreateIdentityOptions{
			CreateTokenWithScopes: []string{"voip"},
			ExpiresInMinutes:      1440,
		},
	)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("failed to create identity")
	}

	client := clientv2.NewWithLogger(key, logger.Default())

	return &_calls{
		host:       host,
		client:     client,
		token:      user.Token,
		validUntil: user.ExpiresOn,
		idc:        &identityClient,
		id:         user.ID,
	}, nil
}
