package vrpboundary

import (
	"context"
	"fmt"
)

// GetValidationStatus retrieves the current public validation state.
func (c *Client) GetValidationStatus(
	ctx context.Context,
	requestID string,
) (*ValidationStatus, error) {

	var status ValidationStatus

	path := fmt.Sprintf(
		"/validation/%s/status",
		requestID,
	)

	if err := c.get(
		ctx,
		path,
		&status,
	); err != nil {
		return nil, err
	}

	return &status, nil
}

// GetValidationVerdict retrieves the final public validation verdict.
func (c *Client) GetValidationVerdict(
	ctx context.Context,
	requestID string,
) (*PublicVerdict, error) {

	var verdict PublicVerdict

	path := fmt.Sprintf(
		"/validation/%s/verdict",
		requestID,
	)

	if err := c.get(
		ctx,
		path,
		&verdict,
	); err != nil {
		return nil, err
	}

	return &verdict, nil
}