package rooms

import "time"

const (
	apiVersion = "2022-02-01"
)

func CreateRoomParticipant(
	id string,
	role Role,
) RoomParticipant {
	return RoomParticipant{
		Role: role,
		CommunicationIdentifier: CommunicationIdentifier{
			RawID: id,
			Id:    id,
		},
	}
}

func RemoveRoomParticipant(
	id string,
) RoomParticipant {
	return RoomParticipant{
		CommunicationIdentifier: CommunicationIdentifier{
			RawID: id,
			Id:    id,
		},
	}
}

type (
	RoomJoinPolicy string
	Role           string
)

const (
	PRESENTER                  Role           = "Presenter"
	ATTENDEE                   Role           = "Attendee"
	CONSUMER                   Role           = "Consumer"
	INVITE_ONLY                RoomJoinPolicy = "InviteOnly"
	COMMUNICATION_SERVICE_USER RoomJoinPolicy = "CommunicationServiceUsers"
)

type CreateRoomOptions struct {
	ValidFrom      time.Time         `json:"validFrom"`
	ValidUntil     time.Time         `json:"validUntil"`
	RoomJoinPolicy RoomJoinPolicy    `json:"roomJoinPolicy"`
	Participants   []RoomParticipant `json:"participants,omitempty"`
}

func (c *CreateRoomOptions) isValid() bool {
	if c.ValidFrom.IsZero() || c.ValidUntil.IsZero() {
		return false
	}
	if c.ValidFrom.After(c.ValidUntil) {
		return false
	}
	if c.RoomJoinPolicy == "" {
		return false
	}
	return true
}

type RoomModel struct {
	Id              string            `json:"id,omitempty"`
	CreatedDateTime time.Time         `json:"createdDateTime,omitempty"`
	ValidFrom       time.Time         `json:"validFrom,omitempty"`
	ValidUntil      time.Time         `json:"validUntil,omitempty"`
	RoomJoinPolicy  RoomJoinPolicy    `json:"roomJoinPolicy,omitempty"`
	Participants    []RoomParticipant `json:"participants,omitempty"`
}

type RoomParticipant struct {
	CommunicationIdentifier CommunicationIdentifier `json:"communicationIdentifier"`
	Role                    Role                    `json:"role,omitempty"`
}

type roomParticipantsUpdate struct {
	Participants []RoomParticipant `json:"participants"`
}

type CommunicationIdentifier struct {
	RawID string `json:"rawId"`
	Id    string `json:"id"`
}

type UpdateRoomOptions CreateRoomOptions

func (c *UpdateRoomOptions) isValid() bool {
	if c.ValidFrom.IsZero() || c.ValidUntil.IsZero() {
		return false
	}
	if c.ValidFrom.After(c.ValidUntil) {
		return false
	}
	if c.RoomJoinPolicy == "" {
		return false
	}
	return true
}
