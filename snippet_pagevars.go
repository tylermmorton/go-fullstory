package fullstory

import "text/template"

var (
	pagevarsSnippetText = `
FS.setVars('page', {
	'pageName': '{{ .PageName }}',
	{{- if .Vars }}
	{{ .Vars -}}
  	{{ end -}}
});
`
	pagevarsTemplate = template.Must(template.New("fs-pagevars").Parse(pagevarsSnippetText))
)

type PageVarsOption func(*pageVars)

func PageVarsWithVar(k string, v interface{}) PageVarsOption {
	return func(p *pageVars) {
		if p.Vars == nil {
			p.Vars = make(Vars)
		}
		p.Vars[k] = v
	}
}

type pageVars struct {
	PageName string
	Vars     Vars
}

func defaultPageVarsOptions(pageName string) *pageVars {
	return &pageVars{
		PageName: pageName,
		Vars:     nil,
	}
}

func PageVarsSnippet(pageName string, opts ...PageVarsOption) (Snippet, error) {
	dotCtx := defaultPageVarsOptions(pageName)
	for _, opt := range opts {
		opt(dotCtx)
	}
	return createSnippetFromTemplate(pagevarsTemplate, dotCtx)
}

func MustPageVarsSnippet(pageName string, opts ...PageVarsOption) Snippet {
	snippet, err := PageVarsSnippet(pageName, opts...)
	if err != nil {
		panic(err)
	}
	return snippet
}
