package vrpboundary

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	DefaultTimeout = 15 * time.Second
	DefaultAgent   = "vrp-public-boundary-go-sdk/0.1.0"
)

// Client provides access to the VRP Public Engineering Boundary.
//
// The client communicates only with public engineering interfaces.
// It never exposes or accesses Protected Runtime functionality.
type Client struct {
	baseURL    string
	httpClient *http.Client
	userAgent  string
}

// NewClient creates a new public boundary client.
func NewClient(baseURL string) *Client {
	baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")

	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		userAgent: DefaultAgent,
	}
}

// SetHTTPClient replaces the default HTTP client.
func (c *Client) SetHTTPClient(client *http.Client) {
	if client != nil {
		c.httpClient = client
	}
}

// SetUserAgent overrides the default SDK user-agent.
func (c *Client) SetUserAgent(agent string) {
	if strings.TrimSpace(agent) != "" {
		c.userAgent = agent
	}
}

func (c *Client) get(
	ctx context.Context,
	path string,
	target any,
) error {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+path,
		nil,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &APIError{
			Code:  ErrorTransportFailure,
			Cause: err,
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {

		apiErr := &APIError{
			HTTPStatus: resp.StatusCode,
		}

		_ = json.Unmarshal(body, apiErr)

		if apiErr.Code == "" {
			apiErr.Code = ErrorUnexpectedResponse
			apiErr.Message = string(body)
		}

		return apiErr
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}