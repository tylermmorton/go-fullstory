package snippet

import (
	_ "embed"
	"net/http"
	"text/template"
)

var (
	//go:embed cookie.tmpl.js
	cookieTmplText string
	cookieTemplate = template.Must(template.New("_fs_cookie").Parse(cookieTmplText))
)

type cookie struct {
	http.Cookie
}

// SessionCookie returns a Snippet that sets a cookie with the given name
// to the value of the FullStory Session ID.
//
// This should be used within the OnReady Snippet.
func SessionCookie(cookieName string) (Snippet, error) {
	return FromTemplate(cookieTemplate, &cookie{
		Cookie: http.Cookie{
			Name:  cookieName,
			Value: "${FS.getCurrentSession()}",

			Path:     "/",
			Secure:   false,
			HttpOnly: false,
		},
	})
}

func MustSessionCookie(cookieName string) Snippet {
	snippet, err := SessionCookie(cookieName)
	if err != nil {
		panic(err)
	}
	return snippet
}
