package fullstory

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

// Config is the FullStory API config.
//
// It can be loaded via github.com/kelseyhightower/envconfig
type Config struct {
	// OrgID is the fullstory organization id
	OrgID string `envconfig:"FULLSTORY_ORG_ID"`
	// APIKey is the fullstory api key
	APIKey string `envconfig:"FULLSTORY_API_KEY"`
}

// Client is the FullStory API client interface
type Client interface {
	CreateEvent(ctx context.Context, req *CreateEventRequest) error
}

type client struct {
	// APIKey is the fullstory api key
	APIKey string
	// Host is the fullstory host
	Host string
	// OrgID is the fullstory organization id
	OrgID string
	// httpClient is the http client
	httpClient *http.Client
}

// NewClient returns a new FullStory API Client
func NewClient(cfg *Config) Client {
	return &client{
		APIKey: cfg.APIKey,
		Host:   "https://api.fullstory.com",
		OrgID:  cfg.OrgID,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) CreateEvent(ctx context.Context, req *CreateEventRequest) error {
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
