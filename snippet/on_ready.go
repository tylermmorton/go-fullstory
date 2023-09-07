package snippet

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed on_ready.tmpl.js
	onReadyTmplText string
	onReadyTemplate = template.Must(template.New("_fs_ready").Parse(onReadyTmplText))
)

type onReady struct {
	Snippets []Snippet
}

func OnReady(snippets ...Snippet) (Snippet, error) {
	return FromTemplate(onReadyTemplate, &onReady{Snippets: snippets})
}

func MustOnReady(snippets ...Snippet) Snippet {
	res, err := OnReady(snippets...)
	if err != nil {
		panic(err)
	}
	return res
}
