package fullstory

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

func createSnippetFromTemplate(template *tt.Template, data interface{}) (Snippet, error) {
	buf := bytes.Buffer{}
	err := template.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return &snippet{buf: buf}, nil
}
