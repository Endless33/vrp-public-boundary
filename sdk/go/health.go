package vrpboundary

import "context"

// GetHealth retrieves the public health endpoint.
//
// The response is intentionally limited to externally observable
// engineering information.
func (c *Client) GetHealth(
	ctx context.Context,
) (*PublicHealth, error) {

	var health PublicHealth

	if err := c.get(
		ctx,
		"/health",
		&health,
	); err != nil {
		return nil, err
	}

	return &health, nil
}

// GetCapabilities retrieves the public capability declaration.
func (c *Client) GetCapabilities(
	ctx context.Context,
) (*PublicCapabilities, error) {

	var capabilities PublicCapabilities

	if err := c.get(
		ctx,
		"/capabilities",
		&capabilities,
	); err != nil {
		return nil, err
	}

	return &capabilities, nil
}