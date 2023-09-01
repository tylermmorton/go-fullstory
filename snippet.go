package fullstory

import (
	"bytes"
	"html/template"
)

var (
	snippetText = /* .html */ `
<script>

</script>
`
	snippetTemplate = template.Must(template.New("fs-snippet").Parse(snippetText))
)

type SnippetOption func(*snippet)

type snippet struct {
	Enabled bool

	OrgID string
}

func defaultSnippetOptions(orgID string) *snippet {
	return &snippet{
		Enabled: true,
		OrgID:   orgID,
	}
}

// SetEnabled turns on/off the rendering of the entire <script> tag. It can be useful to
// control this via an environment variable.
//
//	fullstory.Snippet(os.Getenv("FULLSTORY_ORG_ID"),
//	  fullstory.SetEnabled(os.Getenv("FULLSTORY_ENABLED") == "true"),
//	)
func SetEnabled(value bool) SnippetOption {
	return func(s *snippet) {
		s.Enabled = value
	}
}

// Snippet renders a new FullStory recording snippet within a <script> tag from a Go template.
// One can control how the snippet renders by supplying one or more SnippetOptions
func Snippet(orgID string, opts ...SnippetOption) (template.HTML, error) {
	s := defaultSnippetOptions(orgID)
	for _, opt := range opts {
		opt(s)
	}
	buf := bytes.Buffer{}
	err := snippetTemplate.Execute(&buf, s)
	if err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}

func MustSnippet(orgID string, opts ...SnippetOption) template.HTML {
	snippet, err := Snippet(orgID, opts...)
	if err != nil {
		panic(err)
	}
	return snippet
}
