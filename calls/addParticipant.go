package calls

import (
	"context"
	"fmt"
)

func (c *_calls) AddParticipant(
	ctx context.Context,
	callId string,
	info AddPraticipantRequest,
) error {
	if callId == "" {
		return fmt.Errorf("callId cannot be empty")
	}

	if info.ParticipantToAdd.CommunicationUser.ID == "" {
		return fmt.Errorf("communicationUserId cannot be empty")
	}

	_, err := c.client.Post(
		ctx,
		c.host,
		"/calling/callConnections/"+callId+"/participants:add",
		map[string][]string{
			"api-version": {_apiVersion},
		},
		info,
	)
	if err != nil {
		return err
	}

	return err
}
