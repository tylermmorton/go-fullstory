package fullstory

import "net/http"

func CreateSessionCookieMiddleware(cookieName string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			cookie, err := req.Cookie(cookieName)
			if err != nil {
				h.ServeHTTP(wr, req)
				return
			}

			h.ServeHTTP(wr, req.WithContext(WithSessionID(ctx, cookie.Value)))
		})
	}
}

func CreateSessionHeaderMiddleware(headerName string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			sessionID := req.Header.Get(headerName)
			if len(sessionID) == 0 {
				h.ServeHTTP(wr, req)
				return
			}

			h.ServeHTTP(wr, req.WithContext(WithSessionID(ctx, sessionID)))
		})
	}
}
