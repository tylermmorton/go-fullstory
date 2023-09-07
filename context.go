package fullstory

import (
	"context"
)

type contextKey string

const (
	SessionID contextKey = "session_id"
)

// GetSession returns the current FullStory session from the context, or a blank session if
// none is attached.
//
// The session should be attached to the context via one of the following utilities:
//  1. WithSessionID
//  2. CreateSessionCookieMiddleware
//  3. CreateSessionHeaderMiddleware
func GetSession(ctx context.Context) Session {
	return Session{
		ID: GetSessionID(ctx),
	}
}

// HasSession returns true if the context has a FullStory session attached.
func HasSession(ctx context.Context) bool {
	return len(GetSessionID(ctx)) > 0
}

// GetSessionID returns the session id from the context.
func GetSessionID(ctx context.Context) string {
	if val, ok := ctx.Value(SessionID).(string); !ok {
		return ""
	} else {
		return val
	}
}

// WithSessionID returns a new context with the given session id attached.
func WithSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, SessionID, sessionID)
}
