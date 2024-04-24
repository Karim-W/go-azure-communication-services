package calls

import (
	"context"
	"encoding/json"
	"fmt"
)

// CreateOutBoundCall creates an outbound Azure Communication Services call
func (c *_calls) CreateOutBoundCall(
	ctx context.Context,
	options *CreateOutboundCallOptions,
) (*OutboundCall, error) {
	// check if options is nil
	if options == nil {
		return nil, fmt.Errorf("options cannot be nil")
	}

	// check if targets is nil
	if options.Targets == nil {
		return nil, fmt.Errorf("targets cannot be nil")
	}

	var outboundCall OutboundCall

	reader, err := c.client.Post(
		ctx,
		c.host,
		"/calling/callConnections",
		map[string][]string{
			"api-version": {_apiVersion},
		},
		options,
	)
	if err != nil {
		return nil, err
	}

	// decode the response
	err = json.NewDecoder(reader).Decode(&outboundCall)
	if err != nil {
		return nil, err
	}

	return &outboundCall, nil
}
