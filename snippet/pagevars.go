package snippet

import (
	_ "embed"
	"github.com/tylermmorton/go-fullstory"
	"text/template"
)

var (
	//go:embed pagevars.tmpl.js
	pagevarsTmplText string
	pagevarsTemplate = template.Must(template.New("_fs_pagevars").Parse(pagevarsTmplText))
)

type pageVars struct {
	PageName string
	Vars     fullstory.Vars
}

func PageVarsSnippet(pageName string, vars fullstory.Vars) (Snippet, error) {
	return FromTemplate(pagevarsTemplate, &pageVars{
		PageName: pageName,
		Vars:     vars,
	})
}

func MustPageVarsSnippet(pageName string, vars fullstory.Vars) Snippet {
	snippet, err := PageVarsSnippet(pageName, vars)
	if err != nil {
		panic(err)
	}
	return snippet
}
