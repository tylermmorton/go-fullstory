package snippet

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed on_debug.tmpl.js
	onDebugTmplText string
	onDebugTemplate = template.Must(template.New("_fs_ready").Parse(onDebugTmplText))
)

type onDebug struct {
	Snippets []Snippet
}

func OnDebug(snippets ...Snippet) (Snippet, error) {
	return FromTemplate(onDebugTemplate, &onReady{Snippets: snippets})
}

func MustOnDebug(snippets ...Snippet) Snippet {
	res, err := OnDebug(snippets...)
	if err != nil {
		panic(err)
	}
	return res
}
