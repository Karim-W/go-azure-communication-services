package chat

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	host = ""
	key  = ""
	id   = ""
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
