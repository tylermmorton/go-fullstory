package fullstory

import (
	"text/template"
)

var (
	identifySnippetText = `
{{ if .Enabled -}}
FS.identify('{{ .UserID }}', {
  displayName: '{{ .DisplayName }}',
  email: '{{ .EmailAddress }}',
{{- if .UserVars }}
  {{ .UserVars -}}
{{ end -}}
});
{{- end }}
`
	identifyTemplate = template.Must(template.New("fs-identify").Parse(identifySnippetText))
)

type IdentifyOption func(*identify)

type identify struct {
	UserID       string
	DisplayName  string
	EmailAddress string

	Enabled  bool
	UserVars Vars
}

func defaultIdentifyOptions(userID, displayName, emailAddress string) *identify {
	return &identify{
		Enabled:      true,
		UserID:       userID,
		DisplayName:  displayName,
		EmailAddress: emailAddress,
		UserVars:     nil,
	}
}

// IdentifyEnabled turns on/off the rendering of the entire <script> tag. It can be useful to
// control this via an environment variable.
//
//	fullstory.IdentifySnippet("123456789", "John Doe", "johndoe@example.com", fullstory.IdentifyEnabled(
//	  os.Getenv("FULLSTORY_ENABLED") == "true",
//	))
func IdentifyEnabled(enabled bool) IdentifyOption {
	return func(i *identify) {
		i.Enabled = enabled
	}
}

func IdentifyUserVars(userVars map[string]interface{}) IdentifyOption {
	return func(i *identify) {
		i.UserVars = userVars
	}
}

// IdentifySnippet renders a new FullStory identification recording within a <script> tag from a Go template.
// One can control how the identifier renders by supplying one or more IdentifyOptions
func IdentifySnippet(userID, displayName, emailAddress string, opts ...IdentifyOption) (Snippet, error) {
	dotCtx := defaultIdentifyOptions(userID, displayName, emailAddress)
	for _, opt := range opts {
		opt(dotCtx)
	}
	return createSnippetFromTemplate(identifyTemplate, dotCtx)
}

func MustIdentifySnippet(userID, displayName, emailAddress string, opts ...IdentifyOption) Snippet {
	html, err := IdentifySnippet(userID, displayName, emailAddress, opts...)
	if err != nil {
		panic(err)
	}
	return html
}
