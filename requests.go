package fullstory

import (
	"context"
	"encoding/json"
	"time"
)

type CreateEventRequest struct {
	// Session is the FullStory session to tie this event to
	Session Session `json:"session"`
	// Context is the event context
	Context Context `json:"context"`
	// Name is the name of the event
	Name string `json:"name"`
	// Timestamp is when the event occurred, written out to ISO 8601
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Properties is the custom event payload (max 20)
	Properties map[string]any `json:"properties,omitempty"`
}

func (req *CreateEventRequest) Validate(ctx context.Context) error {
	if len(req.Properties) > 20 {
		return ErrTooManyProperties
	}

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

		// TODO: implement schema generation for properties
		Properties: req.Properties,
	})
}
