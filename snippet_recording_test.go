package fullstory

import (
	"strings"
	"testing"
)

func Test_RecordingSnippet(t *testing.T) {
	testTable := map[string]struct {
		orgID          string
		opts           []RecordingOption
		doContains     []string
		dontContains   []string
		wantErr        bool
		wantErrMessage string
	}{
		"Renders the recording snippet successfully with the default options": {
			orgID:   "123456789",
			opts:    []RecordingOption{},
			wantErr: false,
			doContains: []string{
				`window['_fs_host'] = 'fullstory.com';`,
				`window['_fs_script'] = 'edge.fullstory.com\/s\/fs.js';`,
				`window['_fs_org'] = '123456789';`,
				`window['_fs_namespace'] = 'FS';`,
			},
		},
		"Setting the RecordingEnabled option to false does not render the <script> tag": {
			orgID:   "123456789",
			opts:    []RecordingOption{RecordingEnabled(false)},
			wantErr: false,
			dontContains: []string{
				`window['_fs_host'] = 'fullstory.com';`,
			},
		},
	}

	for name, testCase := range testTable {
		t.Run(name, func(t *testing.T) {
			snippet, err := RecordingSnippet(testCase.orgID, testCase.opts...)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			for _, expect := range testCase.doContains {
				if !strings.Contains(string(snippet.AsJS()), expect) {
					t.Logf("recording: %s", snippet)
					t.Errorf("expected recording to contain %q", expect)
				}
			}

			for _, expect := range testCase.dontContains {
				if strings.Contains(string(snippet.AsJS()), expect) {
					t.Logf("recording: %s", snippet)
					t.Errorf("expected recording to not contain %q", expect)
				}
			}
		})
	}
}
