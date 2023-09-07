package snippet

import (
	"bytes"
	"fmt"
	"html/template"
	tt "text/template"
)

// Snippet represents a piece of JavaScript that can be rendered as either a <script> tag or
// as plain JavaScript within a Go html/template.
type Snippet interface {
	AsJS() template.JS
	AsHTML() template.HTML
}

type snippet struct {
	buf bytes.Buffer
}

func (s *snippet) AsJS() template.JS {
	return template.JS(s.buf.String())
}

func (s *snippet) AsHTML() template.HTML {
	return template.HTML(fmt.Sprintf("<script type=\"text\\javascript\">%s</script>", s.buf.String()))
}

// FromJS creates a new Snippet from a template.JS. This is useful if you want to
// render some JS within OnReady. Be sure to sanitize the JS if its source is from
// an untrusted location.
func FromJS(js template.JS) Snippet {
	return &snippet{buf: *bytes.NewBufferString(string(js))}
}

// FromString creates a new Snippet from a string. Be sure to sanitize the string
// if its source is from an untrusted location because it will be put inside some
// <script> tags.
func FromString(s string) Snippet {
	return &snippet{buf: *bytes.NewBufferString(s)}
}

// FromTemplate creates a new Snippet from a Go text/template. The template will be
// executed with the given data and the result will be returned as a Snippet.
func FromTemplate(template *tt.Template, data interface{}) (Snippet, error) {
	buf := bytes.Buffer{}
	err := template.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return &snippet{buf: buf}, nil
}
