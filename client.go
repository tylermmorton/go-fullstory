package fullstory

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	// OrgID is the fullstory organization id
	OrgID string `env:"FULLSTORY_ORG_ID"`
	// Enabled is whether or not fullstory is enabled
	Enabled bool `env:"FULLSTORY_ENABLED"`
	// APIKey is the fullstory api key
	APIKey string `env:"FULLSTORY_API_KEY"`
}

type Client struct {
	// APIKey is the fullstory api key
	APIKey string
	// Enabled is whether or not fullstory is enabled
	Enabled bool
	// Host is the fullstory host
	Host string
	// OrgID is the fullstory organization id
	OrgID string

	// httpClient is the http client
	httpClient *http.Client
}

// NewClient returns a new fullstory client
func NewClient(cfg *Config) *Client {
	return &Client{
		APIKey:  cfg.APIKey,
		Enabled: cfg.Enabled,
		Host:    "https://api.fullstory.com",
		OrgID:   cfg.OrgID,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) PostEvent(ctx context.Context, req *CreateEventRequest) error {
	byt, err := req.MarshalJSON()
	if err != nil {
		return err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/v2/events", c.Host), bytes.NewReader(byt))
	if err != nil {
		return err
	}

	r.Header.Add("Authorization", "Basic "+c.APIKey)
	r.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create new event: %s", res.Status)
	}

	return nil
}
