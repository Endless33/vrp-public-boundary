package vrpboundary

import (
	"context"
	"fmt"
)

// GetEvidenceReference retrieves the public reference to an evidence artifact.
//
// The returned object contains only publicly observable metadata.
// It never exposes Protected Runtime implementation or internal state.
func (c *Client) GetEvidenceReference(
	ctx context.Context,
	evidenceID string,
) (*EvidenceReference, error) {

	var evidence EvidenceReference

	path := fmt.Sprintf(
		"/evidence/%s",
		evidenceID,
	)

	if err := c.get(
		ctx,
		path,
		&evidence,
	); err != nil {
		return nil, err
	}

	return &evidence, nil
}