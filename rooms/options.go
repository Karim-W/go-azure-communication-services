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
	role Role,
) RoomParticipant {
	return RoomParticipant{
		Role: role,
		CommunicationIdentifier: CommunicationIdentifier{
			RawID: id,
			CommunicationUser: CommunicationUser{
				ID: id,
			},
		},
	}
}

type Role string

const (
	PRESENTER Role = "Presenter"
	ATTENDEE  Role = "Attendee"
	CONSUMER  Role = "Consumer"
)

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
	CommunicationIdentifier CommunicationIdentifier `json:"communicationIdentifier"`
	Role                    Role                    `json:"role,omitempty"`
}

type CommunicationIdentifier struct {
	RawID             string            `json:"rawId"`
	CommunicationUser CommunicationUser `json:"communicationUser"`
}

type CommunicationUser struct {
	ID string `json:"id"`
}

// type CommunicationIdentifierModel struct {
// 	RawId             string                           `json:"rawId"`
// 	CommunicationUser CommunicationUserIdentifierModel `json:"communicationUser"`
// }

type CommunicationUserIdentifierModel struct {
	Id string `json:"communicationUserId"`
}

type UpdateRoomOptions CreateRoomOptions
