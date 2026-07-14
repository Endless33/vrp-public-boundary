package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://pilot.example.invalid/vrp/api/v1"
	maxBodyBytes   = 1 << 20
)

type PublicHealth struct {
	Schema          string    `json:"schema"`
	Status          string    `json:"status"`
	ObservedAt      time.Time `json:"observed_at"`
	BoundaryVersion string    `json:"boundary_version"`
}

type PublicCapabilities struct {
	Schema          string   `json:"schema"`
	BoundaryVersion string   `json:"boundary_version"`
	Capabilities    []string `json:"capabilities"`
}

type PublicError struct {
	Schema    string `json:"schema"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) (*Client, error) {
	if baseURL == "" {
		return nil, errors.New("base URL is required")
	}

	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

func (c *Client) GetHealth(ctx context.Context) (PublicHealth, error) {
	var health PublicHealth

	if err := c.getJSON(ctx, "/health", &health); err != nil {
		return PublicHealth{}, err
	}

	if health.Schema != "vrp.health/v1" {
		return PublicHealth{}, fmt.Errorf(
			"unexpected health schema: %q",
			health.Schema,
		)
	}

	return health, nil
}

func (c *Client) GetCapabilities(
	ctx context.Context,
) (PublicCapabilities, error) {
	var capabilities PublicCapabilities

	if err := c.getJSON(
		ctx,
		"/capabilities",
		&capabilities,
	); err != nil {
		return PublicCapabilities{}, err
	}

	if capabilities.Schema != "vrp.capabilities/v1" {
		return PublicCapabilities{}, fmt.Errorf(
			"unexpected capabilities schema: %q",
			capabilities.Schema,
		)
	}

	return capabilities, nil
}

func (c *Client) getJSON(
	ctx context.Context,
	path string,
	target any,
) error {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+path,
		nil,
	)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	request.Header.Set("Accept", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer response.Body.Close()

	body := io.LimitReader(response.Body, maxBodyBytes)

	if response.StatusCode < 200 || response.StatusCode > 299 {
		var publicError PublicError

		if decodeErr := json.NewDecoder(body).Decode(
			&publicError,
		); decodeErr != nil {
			return fmt.Errorf(
				"public API returned HTTP %d",
				response.StatusCode,
			)
		}

		return fmt.Errorf(
			"public API error: code=%s message=%s",
			publicError.Code,
			publicError.Message,
		)
	}

	if err := json.NewDecoder(body).Decode(target); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}

func main() {
	client, err := NewClient(defaultBaseURL)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer cancel()

	health, err := client.GetHealth(ctx)
	if err != nil {
		fmt.Printf("health request failed: %v\n", err)
		return
	}

	fmt.Printf(
		"schema=%s status=%s boundary_version=%s observed_at=%s\n",
		health.Schema,
		health.Status,
		health.BoundaryVersion,
		health.ObservedAt.UTC().Format(time.RFC3339),
	)

	capabilities, err := client.GetCapabilities(ctx)
	if err != nil {
		fmt.Printf("capabilities request failed: %v\n", err)
		return
	}

	fmt.Printf(
		"schema=%s boundary_version=%s capabilities=%v\n",
		capabilities.Schema,
		capabilities.BoundaryVersion,
		capabilities.Capabilities,
	)
}