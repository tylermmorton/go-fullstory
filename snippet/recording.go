package snippet

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed recording.tmpl.js
	recordingTmplText string
	recordingTemplate = template.Must(template.New("_fs_recording").Parse(recordingTmplText))
)

type RecordingOption func(*recording)

type recording struct {
	Host      string
	Script    string
	OrgID     string
	Namespace string
}

func defaultRecordingOptions(orgID string) *recording {
	return &recording{
		Host:      "fullstory.com",
		Script:    "edge.fullstory.com/s/fs.js",
		OrgID:     orgID,
		Namespace: "FS",
	}
}

// RecordingHost sets the host of the FullStory recording snippet.
func RecordingHost(value string) RecordingOption {
	return func(s *recording) {
		s.Host = value
	}
}

// RecordingScript sets the script URL of the FullStory recording snippet.
func RecordingScript(value string) RecordingOption {
	return func(s *recording) {
		s.Script = value
	}
}

// RecordingNamespace sets the namespace of the FullStory recording snippet.
func RecordingNamespace(value string) RecordingOption {
	return func(s *recording) {
		s.Namespace = value
	}
}

// RecordingSnippet renders a new FullStory recording snippet within a <script> tag from a Go template.
// One can control how the recording renders by supplying one or more RecordingOption
func RecordingSnippet(orgID string, opts ...RecordingOption) (Snippet, error) {
	dotCtx := defaultRecordingOptions(orgID)
	for _, opt := range opts {
		opt(dotCtx)
	}
	return FromTemplate(recordingTemplate, dotCtx)
}

func MustRecordingSnippet(orgID string, opts ...RecordingOption) Snippet {
	snippet, err := RecordingSnippet(orgID, opts...)
	if err != nil {
		panic(err)
	}
	return snippet
}
