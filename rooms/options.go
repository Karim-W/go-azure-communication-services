package rooms

import "time"

const (
	apiVersion = "2022-02-01"
)

type UserTuple struct {
	AcsId string `json:"userId"`
	Role  string `json:"role"`
}

func CreateRoomParticipant(
	id string,
	role string,
) RoomParticipant {
	return RoomParticipant{
		Role:                    role,
		CommunicationIdentifier: id,
	}
}

type CreateRoomOptions struct {
	ValidFrom      time.Time         `json:"validFrom"`
	ValidUntil     time.Time         `json:"validUntil"`
	RoomJoinPolicy string            `json:"roomJoinPolicy"`
	Participants   []RoomParticipant `json:"participants,omitempty"`
}

type RoomModel struct {
	Id              string            `json:"id,omitempty"`
	CreatedDateTime time.Time         `json:"createdDateTime,omitempty"`
	ValidFrom       time.Time         `json:"validFrom,omitempty"`
	ValidUntil      time.Time         `json:"validUntil,omitempty"`
	RoomJoinPolicy  string            `json:"roomJoinPolicy,omitempty"`
	Participants    []RoomParticipant `json:"participants,omitempty"`
}

type RoomParticipant struct {
	CommunicationIdentifier string `json:"communicationIdentifier"`
	Role                    string `json:"role"`
}

// type CommunicationIdentifierModel struct {
// 	RawId             string                           `json:"rawId"`
// 	CommunicationUser CommunicationUserIdentifierModel `json:"communicationUser"`
// }

type CommunicationUserIdentifierModel struct {
	Id string `json:"communicationUserId"`
}

type UpdateRoomOptions CreateRoomOptions
