package snippet

import (
	_ "embed"
	"github.com/tylermmorton/go-fullstory"
	"text/template"
)

var (
	//go:embed identify.tmpl.js
	identifySnippetText string
	identifyTemplate    = template.Must(template.New("_fs_identify").Parse(identifySnippetText))
)

type identify struct {
	UserID       string
	DisplayName  string
	EmailAddress string
	UserVars     fullstory.Vars
}

// IdentifySnippet renders a new FullStory identification recording within a <script> tag from a Go template.
// One can control how the identifier renders by supplying one or more IdentifyOptions
func IdentifySnippet(userID, displayName, emailAddress string, userVars fullstory.Vars) (Snippet, error) {
	return FromTemplate(identifyTemplate, &identify{
		UserID:       userID,
		DisplayName:  displayName,
		EmailAddress: emailAddress,
		UserVars:     userVars,
	})
}

func MustIdentifySnippet(userID, displayName, emailAddress string, userVars fullstory.Vars) Snippet {
	html, err := IdentifySnippet(userID, displayName, emailAddress, userVars)
	if err != nil {
		panic(err)
	}
	return html
}
