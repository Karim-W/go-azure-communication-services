package calls

import (
	"fmt"
)

const _apiVersion = "2023-10-15"

type CreateOutboundCallOptions struct {
	Targets     []Target `json:"targets"`
	CallbackURI string   `json:"callbackUri"`
}

type Target struct {
	Kind              string            `json:"kind"`
	CommunicationUser CommunicationUser `json:"communicationUser"`
}

var (
	ERR_UNAUTHORIZED      = fmt.Errorf("unauthorized")
	ERR_EXPIRED_TOKEN     = fmt.Errorf("token expired")
	ERR_NO_TOKEN_PROVIDED = fmt.Errorf("no token provided")
)

type OutboundCall struct {
	CallConnectionID    string   `json:"callConnectionId"`
	ServerCallID        string   `json:"serverCallId"`
	Source              Source   `json:"source"`
	Targets             []Source `json:"targets"`
	CallConnectionState string   `json:"callConnectionState"`
	CallbackURI         string   `json:"callbackUri"`
}

type Source struct {
	Kind              string            `json:"kind"`
	CommunicationUser CommunicationUser `json:"communicationUser"`
}

type AddPraticipantRequest struct {
	ParticipantToAdd     ParticipantToAdd     `json:"participantToAdd"`
	OperationContext     string               `json:"operationContext"`
	SourceCallerIDNumber SourceCallerIDNumber `json:"sourceCallerIdNumber"`
	OperationCallbackURI string               `json:"operationCallbackUri"`
}

type ParticipantToAdd struct {
	Kind              string            `json:"kind"`
	CommunicationUser CommunicationUser `json:"communicationUser"`
}

type CommunicationUser struct {
	ID string `json:"id"`
}

type SourceCallerIDNumber struct {
	Value string `json:"value"`
}
