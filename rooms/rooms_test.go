package rooms

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	host   = ""
	key    = ""
	id     = ""
	roomid = ""
)

func precheck() {
	host = os.Getenv("ACS_HOST")
	key = os.Getenv("ACS_KEY")
	id = os.Getenv("ACS_USER_ID")
	roomid = os.Getenv("ACS_ROOM_ID")
}

func TestCreateRoom(t *testing.T) {
	precheck()
	if host == "" || key == "" || id == "" {
		t.Skip("no host, key or id")
	}

	client := New(host, key)
	room, err := client.CreateRoom(
		context.TODO(),
		&CreateRoomOptions{
			ValidFrom:      time.Now(),
			ValidUntil:     time.Now().Add(time.Hour),
			RoomJoinPolicy: INVITE_ONLY,
			Participants: []RoomParticipant{
				CreateRoomParticipant(id, PRESENTER),
			},
		},
	)

	t.Log("created room", room.Id)

	assert.Nil(t, err)
	assert.NotNil(t, room)
}

func TestGetRoom(t *testing.T) {
	precheck()
	if host == "" || key == "" || roomid == "" {
		t.Skip("no host, key or roomid")
	}

	client := New(host, key)
	room, err := client.GetRoom(
		context.TODO(),
		roomid,
	)
	assert.Nil(t, err)
	assert.NotNil(t, room)
}

func TestUpdateRoom(t *testing.T) {
	precheck()
	if host == "" || key == "" || roomid == "" || id == "" {
		t.Skip("no host, key, roomid or id")
	}

	client := New(host, key)
	room, err := client.UpdateRoom(
		context.TODO(),
		roomid,
		&UpdateRoomOptions{
			ValidFrom:      time.Now(),
			ValidUntil:     time.Now().Add(time.Hour),
			RoomJoinPolicy: INVITE_ONLY,
			Participants: []RoomParticipant{
				CreateRoomParticipant(id, PRESENTER),
			},
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, room)
}

func TestDeleteRoom(t *testing.T) {
	precheck()
	if host == "" || key == "" || roomid == "" {
		t.Skip("no host, key or roomid")
	}

	client := New(host, key)
	err := client.DeleteRoom(
		context.TODO(),
		roomid,
	)
	assert.Nil(t, err)
}

func TestAddParticipant(t *testing.T) {
	precheck()
	if host == "" || key == "" || roomid == "" || id == "" {
		t.Skip("no host, key, roomid or id")
	}

	client := New(host, key)
	room, err := client.AddParticipants(
		context.TODO(),
		roomid,
		CreateRoomParticipant(id, PRESENTER),
	)
	assert.Nil(t, err)
	assert.NotNil(t, room)
}

func TestGetParticipants(t *testing.T) {
	precheck()
	if host == "" || key == "" || roomid == "" {
		t.Skip("no host, key or roomid")
	}
	client := New(host, key)
	room, err := client.GetParticipants(
		context.TODO(),
		roomid,
	)
	assert.Nil(t, err)
	assert.NotNil(t, room)
}

func TestUpdateParticipants(t *testing.T) {
	precheck()
	if host == "" || key == "" || roomid == "" || id == "" {
		t.Skip("no host, key, roomid or id")
	}

	client := New(host, key)
	room, err := client.UpdateParticipants(
		context.TODO(),
		roomid,
		CreateRoomParticipant(id, ATTENDEE),
	)
	assert.Nil(t, err)
	assert.NotNil(t, room)
}

func TestRemoveParticipant(t *testing.T) {
	precheck()
	if host == "" || key == "" || roomid == "" || id == "" {
		t.Skip("no host, key, roomid or id")
	}

	client := New(host, key)
	room, err := client.RemoveParticipants(
		context.TODO(),
		roomid,
		RemoveRoomParticipant(id),
	)
	assert.Nil(t, err)
	assert.NotNil(t, room)
}
