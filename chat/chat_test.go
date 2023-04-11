package chat

import (
	"context"
	"testing"
	"time"

	"github.com/karim-w/go-azure-communication-services/identity"
	"github.com/stretchr/testify/assert"
)

var (
	host     = ""
	key      = ""
	id       = ""
	id2      = ""
	token    = ""
	threadId = ""
)

func TestCreateChatThread(t *testing.T) {
	client, err := New(host, key)
	assert.Nil(t, err)
	thread, err := client.CreateChatThread(
		context.Background(),
		"test",
		ChatUser{
			ID:          id,
			DisplayName: "test",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, thread)
	assert.NotEmpty(t, thread.ChatThread.ID)
}

func TestDeleteChatThread(t *testing.T) {
	client, err := New(host, key)
	assert.Nil(t, err)
	thread, err := client.CreateChatThread(
		context.Background(),
		"test",
		ChatUser{
			ID:          id,
			DisplayName: "test",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, thread)
	assert.NotEmpty(t, thread.ChatThread.ID)
	err = client.DeleteChatThread(
		context.Background(),
		thread.ChatThread.ID,
	)
	assert.Nil(t, err)
}

func TestAddChatParticipant(t *testing.T) {
	client, err := NewWithToken(host, token, time.Now().Add(time.Hour))
	assert.Nil(t, err)
	thread, err := client.CreateChatThread(
		context.Background(),
		"test",
		ChatUser{
			ID:          id,
			DisplayName: "test",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, thread)
	assert.NotEmpty(t, thread.ChatThread.ID)
	err = client.AddChatParticipants(
		context.Background(),
		thread.ChatThread.ID,
		ChatUser{
			ID:          id2,
			DisplayName: "test2",
		},
	)
	assert.Nil(t, err)
}

func TestRemoveChatParticipant(t *testing.T) {
	client, err := NewWithToken(host, token, time.Now().Add(time.Hour))
	assert.Nil(t, err)
	thread, err := client.CreateChatThread(
		context.Background(),
		"test",
		ChatUser{
			ID:          id,
			DisplayName: "test",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, thread)
	assert.NotEmpty(t, thread.ChatThread.ID)
	err = client.AddChatParticipants(
		context.Background(),
		thread.ChatThread.ID,
		ChatUser{
			ID:          id2,
			DisplayName: "test2",
		},
	)
	assert.Nil(t, err)
	err = client.RemoveChatParticipant(
		context.Background(),
		thread.ChatThread.ID,
		identity.CommunicationIdentifier{
			RawID: id2,
			CommunicationUser: identity.CommunicationUser{
				ID: id2,
			},
		},
	)
	assert.Nil(t, err)
}
