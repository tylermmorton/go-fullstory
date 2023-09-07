> 👋🏻 This package is currently under development. Please check back soon for more updates!

# FullStory Go SDK

This package is an unofficial Go SDK for [FullStory](https://www.fullstory.com?ref=go-fullstory). It provides some useful tools for quickly integrating FullStory into your Go web application.

## Getting Started

Getting started with FullStory is easy. Sign up for a free account at [fullstory.com](https://www.fullstory.com/signup) -- After, you'll be able to find your [Org ID](https://help.fullstory.com/hc/en-us/articles/360047075853-How-do-I-find-my-FullStory-Org-Id-?ref=go-fullstory) and [API key](https://help.fullstory.com/hc/en-us/articles/360020624834-Where-can-I-find-my-API-key-?ref=go-fullstory) in the Settings page.

## Usage

### Snippets

This package provides a series of "snippets" that you'll want to include in your web pages to enable recording and playback of user sessions in FullStory. This package provides a simple interface for generating these snippets for use in Go html/templates.

In code, a `Snippet` is simply some JavaScript text wrapped in a type that implements the `Snippet` interface. One can use this interface to render the same snippet as both a JS expression or wrapped in `<script>` tags.

```go
type Snippet interface {
	AsJS() template.JS
	AsHTML() template.HTML
}
```

#### Recording Snippet

**[FullStory Documentation](https://help.fullstory.com/hc/en-us/articles/360020623514-Installing-the-FullStory-Script?ref=go-fullstory)**

The recording script is key to FullStory's functionality and enables tracking of the activity in your web application. It is required to use FullStory and should be included in the `<head>` of your HTML pages.

This package provides a helper function for generating a FullStory recording script that can be used in your Go templates:

```go
package main

import (
	"github.com/tylermmorton/go-fullstory/snippet"
	"html/template"
	"net/http"
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
			"FullStorySnippet": snippet.MustRecordingSnippet("YOUR_ORG_ID").AsHTML(),
		})
	})
	http.ListenAndServe(":8080", nil)
}
```

#### Identify Snippet

**[FullStory Documentation](https://help.fullstory.com/hc/en-us/articles/360020828113-FS-identify-Identifying-users?ref=go-fullstory)**

The identify snippet is used to associate a user with a FullStory session. This snippet should be included on every page of your web application where a user has been authenticated.

Additionally, you can use the identify snippet to set user variables that will be indexed and available in FullStory. This is useful for associating a user's name, email, or other information with a FullStory session.

```go
func IdentifySnippet(userID, displayName, emailAddress string, userVars fullstory.Vars) (Snippet, error)
```

And a `Must` variant that will panic if an error occurs:
```go
func MustIdentifySnippet(userID, displayName, emailAddress string, userVars fullstory.Vars) Snippet
```

#### PageVars Snippet

**[FullStory Documentation](https://help.fullstory.com/hc/en-us/articles/1500004101581-FS-setVars-API-Sending-custom-page-data-to-FullStory?ref=go-fullstory)**

The pagevars snippet is used to set page variables that will be indexed and available in FullStory. This is useful for associating information about the page a user is on with a FullStory session.

```go
func PageVarsSnippet(pageVars fullstory.Vars) (Snippet, error)
```

And a `Must` variant that will panic if an error occurs:
```go
func MustPageVarsSnippet(pageVars fullstory.Vars) Snippet
```

### HTTP API Client

This package also provides an implementation of the FullStory Server API. One can create an API client like so:

```go
package main

import (
	"github.com/tylermmorton/go-fullstory"
	"os"
)

func main() {
	fs := fullstory.NewClient(&fullstory.Config{
		OrgID:  os.Getenv("FULLSTORY_ORG_ID"),
		APIKey: os.Getenv("FULLSTORY_API_KEY"),
	})
}
```
