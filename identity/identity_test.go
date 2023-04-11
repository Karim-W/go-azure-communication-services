package identity

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	host = ""
	key  = ""
)

func TestCreateIdentityUser(t *testing.T) {
	identity := New(host, key)
	user, err := identity.CreateIdentity(
		context.Background(),
		&CreateIdentityOptions{
			CreateTokenWithScopes: []string{"chat", "voip"},
			ExpiresInMinutes:      60,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
