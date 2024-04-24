package calls

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	host     = os.Getenv("ACS_HOST")
	key      = os.Getenv("ACS_KEY")
	id       = os.Getenv("ACS_USER_ID")
	id2      = os.Getenv("ACS_USER_ID2")
	token    = ""
	threadId = ""
)

func TestTokenSetter(t *testing.T) {
	c := _calls{}
	setter := func() (string, error) {
		return "test", nil
	}
	c.SetTokenFetcher(setter)
	assert.NotNil(t, c.tokenFetcher)
}

func TestTokenGetter(t *testing.T) {
	c := _calls{}
	setter := func() (string, error) {
		return "test", nil
	}
	c.SetTokenFetcher(setter)
	token, err := c.GetToken()
	assert.Nil(t, err)
	assert.Equal(t, "test", token)
}

func TestTokenGetterWithError(t *testing.T) {
	c := _calls{}
	setter := func() (string, error) {
		return "", fmt.Errorf("test")
	}
	c.SetTokenFetcher(setter)
	token, err := c.GetToken()
	assert.NotNil(t, err)
	assert.Equal(t, "", token)
}

func TestCreateOutboundCall(t *testing.T) {
	client, err := New(host, key)
	assert.Nil(t, err)

	c, err := client.CreateOutBoundCall(
		context.Background(),
		&CreateOutboundCallOptions{
			Targets: []Target{
				{
					Kind: "communicationUser",
					CommunicationUser: CommunicationUser{
						ID: id,
					},
				},
			},
			CallbackURI: "https://example.com",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, c)

	byts, err := json.Marshal(c)
	assert.Nil(t, err)
	fmt.Println(string(byts))
}

func TestAddParticipantToCalls(t *testing.T) {
	client, err := New(host, key)
	assert.Nil(t, err)

	c, err := client.CreateOutBoundCall(
		context.Background(),
		&CreateOutboundCallOptions{
			Targets: []Target{
				{
					Kind: "communicationUser",
					CommunicationUser: CommunicationUser{
						ID: id,
					},
				},
			},
			CallbackURI: "https://example.com",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.NotEmpty(t, c.CallConnectionID)

	byts, err := json.Marshal(c)
	assert.Nil(t, err)
	fmt.Println(string(byts))

	time.Sleep(15 * time.Second)

	err = client.AddParticipant(
		context.Background(),
		c.CallConnectionID,
		AddPraticipantRequest{
			ParticipantToAdd: ParticipantToAdd{
				Kind: "communicationUser",
				CommunicationUser: CommunicationUser{
					ID: id2,
				},
			},
			OperationContext: "test",
			SourceCallerIDNumber: SourceCallerIDNumber{
				Value: "test",
			},
			OperationCallbackURI: "https://example.com",
		},
	)
	assert.Nil(t, err)
}
