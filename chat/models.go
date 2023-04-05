package chat

import (
	"github.com/karim-w/go-azure-communication-services/identity"
)

type CreateChatThread struct {
	Topic        string        `json:"topic"`
	Participants []Participant `json:"participants"`
}

const _apiVersion = "2021-09-07"

type Participant struct {
	CommunicationIdentifier identity.CommunicationIdentifier `json:"communicationIdentifier"`
	DisplayName             string                           `json:"displayName"`
}

type ChatUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type CreateChatThreadResponse struct {
	ChatThread          ChatThread           `json:"chatThread"`
	InvalidParticipants []InvalidParticipant `json:"invalidParticipants"`
}

type ChatThread struct {
	ID                               string                           `json:"id"`
	Topic                            string                           `json:"topic"`
	CreatedOn                        string                           `json:"createdOn"`
	CreatedByCommunicationIdentifier CreatedByCommunicationIdentifier `json:"createdByCommunicationIdentifier"`
}

type CreatedByCommunicationIdentifier struct {
	RawID             string            `json:"rawId"`
	CommunicationUser CommunicationUser `json:"communicationUser"`
}

type CommunicationUser struct {
	ID string `json:"id"`
}

type InvalidParticipant struct {
	Target  string `json:"target"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
