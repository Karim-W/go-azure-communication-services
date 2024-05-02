package identity

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	host = os.Getenv("ACS_HOST")
	key  = os.Getenv("ACS_KEY")
)

func TestCreateIdentityUser(t *testing.T) {
	if host == "" {
		t.Skip("no host")
	}

	if key == "" {
		t.Skip("no key")
	}

	identity := NewWithLogger(host, key, nil)

	user, err := identity.CreateIdentity(
		context.Background(),
		&CreateIdentityOptions{
			CreateTokenWithScopes: []string{"chat", "voip"},
			ExpiresInMinutes:      60,
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Token)
	assert.NotEmpty(t, user.ExpiresOn)
}

func TestIssueAccessToken(t *testing.T) {
	if host == "" {
		t.Skip("no host")
	}

	if key == "" {
		t.Skip("no key")
	}

	identity := NewWithLogger(host, key, nil)

	user, err := identity.CreateIdentity(
		context.Background(),
		&CreateIdentityOptions{
			CreateTokenWithScopes: []string{"chat", "voip"},
			ExpiresInMinutes:      60,
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	user, err = identity.IssueAccessToken(
		context.Background(),
		user.ID,
		&IssueTokenOptions{
			ExpiresInMinutes: 60,
			Scopes:           []string{"chat", "voip"},
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Token)
	assert.NotEmpty(t, user.ExpiresOn)
}

func TestDeleteIdentity(t *testing.T) {
	if host == "" {
		t.Skip("no host")
	}

	if key == "" {
		t.Skip("no key")
	}

	identity := NewWithLogger(host, key, nil)

	user, err := identity.CreateIdentity(
		context.Background(),
		&CreateIdentityOptions{
			CreateTokenWithScopes: []string{"chat", "voip"},
			ExpiresInMinutes:      60,
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	err = identity.DeleteIdentity(
		context.Background(),
		user.ID,
	)

	assert.Nil(t, err)
}

func TestRevokeAccessToken(t *testing.T) {
	if host == "" {
		t.Skip("no host")
	}

	if key == "" {
		t.Skip("no key")
	}

	identity := NewWithLogger(host, key, nil)

	user, err := identity.CreateIdentity(
		context.Background(),
		&CreateIdentityOptions{
			CreateTokenWithScopes: []string{"chat", "voip"},
			ExpiresInMinutes:      60,
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	err = identity.RevokeAccessToken(
		context.Background(),
		user.ID,
	)

	assert.Nil(t, err)
}
