package chat

import (
	"context"
	"fmt"
	"testing"
	"time"

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

func TestTokenSetter(t *testing.T) {
	c := _chat{}
	setter := func() (string, error) {
		return "test", nil
	}
	c.SetTokenFetcher(setter)
	assert.NotNil(t, c.tokenFetcher)
}

func TestTokenGetter(t *testing.T) {
	c := _chat{}
	setter := func() (string, error) {
		return "test", nil
	}
	c.SetTokenFetcher(setter)
	token, err := c.GetToken()
	assert.Nil(t, err)
	assert.Equal(t, "test", token)
}

func TestTokenGetterWithError(t *testing.T) {
	c := _chat{}
	setter := func() (string, error) {
		return "", fmt.Errorf("test")
	}
	c.SetTokenFetcher(setter)
	token, err := c.GetToken()
	assert.NotNil(t, err)
	assert.Equal(t, "", token)
}

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
		id,
	)
	assert.Nil(t, err)
}
