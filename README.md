# FullStory Go SDK

This package is an unofficial Go SDK for [FullStory](https://www.fullstory.com). It provides some basic tools for quickly integrating FullStory into your Go web application.

## Snippets

FullStory provides a series of 'snippets' that you'll want to include in your web pages to enable recording and playback of user sessions. This package provides a simple interface for generating these snippets for use in Go html/templates.

In code, a `Snippet` is simply some JavaScript text wrapped in a type that implements the `Snippet` interface. One can use this interface to convert the same snippet to both its raw JS text or wrapped in <script> tags.
```go
type Snippet interface {
	AsJS() template.JS
	AsHTML() template.HTML
}
```

### Recording Snippet

The FullStory Recording Snippet enables FullStory to record the activity in your web application. It should be included in the `<head>` of your HTML pages.

This package provides a helper function for generating a FullStory Recording Snippet that can be used in your Go templates:

```go
package main

import (
	"github.com/tylermmorton/go-fullstory"
	"html/template"
	"net/http"

	"github.com/fullstorydev/fullstory-go"
)

var tmpl = template.Must(template.New("index.html").Parse(`
<html>
    <head>
        {{ .FullStorySnippet }}
    </head>
    <body>
        <h1>Hello, world!</h1>
    </body>
</html>
`))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, map[string]interface{}{
			"FullStorySnippet": fullstory.MustRecordingSnippet("YOUR_ORG_ID").AsHTML(),
		})
	})
	http.ListenAndServe(":8080", nil)
}
```