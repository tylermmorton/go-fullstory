package fullstory

import (
	"context"
	"encoding/json"
	"time"
)

type Session struct {
	// ID is the fullstory session id
	ID string `json:"id"`
}

type Browser struct {
	URL             string `json:"url,omitempty"`
	UserAgent       string `json:"user_agent,omitempty"`
	InitialRefferer string `json:"initial_referrer,omitempty"`
}

type Mobile struct {
	AppID        string `json:"app_id,omitempty"`
	AppVersion   string `json:"app_version,omitempty"`
	AppName      string `json:"app_name,omitempty"`
	BuildVariant string `json:"build_variant,omitempty"`
}

type Device struct {
	Manufacturer   string `json:"manufacturer,omitempty"`
	Model          string `json:"model,omitempty"`
	ScreenWidth    int32  `json:"screen_width,omitempty"`
	ScreenHeight   int32  `json:"screen_height,omitempty"`
	ViewportWidth  int32  `json:"viewport_width,omitempty"`
	ViewportHeight int32  `json:"viewport_height,omitempty"`
}

type Location struct{}

type Context struct {
	Browser  Browser  `json:"browser,omitempty"`
	Mobile   Mobile   `json:"mobile,omitempty"`
	Device   Device   `json:"device,omitempty"`
	Location Location `json:"location,omitempty"`
}

type CreateEventRequest struct {
	Session Session `json:"session"`
	Context Context `json:"context"`

	// Name is the name of the event
	Name string `json:"name"`
	// ISO 8601 string optional
	Timestamp time.Time `json:"timestamp,omitempty"`
	// The custom event payload
	Properties map[string]any `json:"properties,omitempty"`
}

func (req *CreateEventRequest) Validate(ctx context.Context) error {
	return nil
}


func (req *CreateEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Session Session `json:"session"`
		Context Context `json:"context"`

		Name      string `json:"name"`
		Timestamp string `json:"timestamp,omitempty"`

		Properties map[string]any `json:"properties,omitempty"`
	}{
		Session: req.Session,
		Context: req.Context,

		Name:      req.Name,
		Timestamp: req.Timestamp.Format(time.RFC3339), // ISO 8601

		Properties: req.Properties,
	})
}
