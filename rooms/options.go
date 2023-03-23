package rooms

import "time"

const (
	apiVersion = "2022-02-01"
)

type CreateRoomOptions struct {
	ValidFrom      time.Time          `json:"validFrom,omitempty"`
	ValidUntil     time.Time          `json:"validUntil,omitempty"`
	RoomJoinPolicy string             `json:"roomJoinPolicy,omitempty"`
	Participants   []RoomParticipants `json:"participants,omitempty"`
}

type RoomModel struct {
	Id              string             `json:"id,omitempty"`
	CreatedDateTime time.Time          `json:"createdDateTime,omitempty"`
	ValidFrom       time.Time          `json:"validFrom,omitempty"`
	ValidUntil      time.Time          `json:"validUntil,omitempty"`
	RoomJoinPolicy  string             `json:"roomJoinPolicy,omitempty"`
	Participants    []RoomParticipants `json:"participants,omitempty"`
}

type RoomParticipants struct {
	CommunicationIdentifier CommunicationIdentifierModel `json:"communicationIdentifier,omitempty"`
	Role                    string                       `json:"role,omitempty"`
}

type CommunicationIdentifierModel struct {
	RawId             string                           `json:"rawId,omitempty"`
	CommunicationUser CommunicationUserIdentifierModel `json:"communicationUser,omitempty"`
}

type CommunicationUserIdentifierModel struct {
	Id string `json:"id,omitempty"`
}

type UpdateRoomOptions CreateRoomOptions
