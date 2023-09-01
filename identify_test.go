package fullstory

import (
	"strings"
	"testing"
)

func Test_Identify(t *testing.T) {
	testTable := map[string]struct {
		userID         string
		displayName    string
		emailAddress   string
		opts           []IdentifyOption
		doContains     []string
		dontContains   []string
		wantErr        bool
		wantErrMessage string
	}{
		"Renders the identify recording successfully with the default options": {
			userID:       "123456789",
			displayName:  "John Doe",
			emailAddress: "johndoe@example.com",
			opts:         []IdentifyOption{},
			wantErr:      false,
			doContains: []string{
				"FS.identify('123456789', {",
				"displayName: 'John Doe',",
				"email: 'johndoe@example.com',",
			},
		},
		"Setting the IdentifyEnabled option to false does not render the identify recording": {
			userID:       "123456789",
			displayName:  "John Doe",
			emailAddress: "",
			opts:         []IdentifyOption{IdentifyEnabled(false)},
			wantErr:      false,
			dontContains: []string{
				"FS.identify('123456789', {",
				"displayName: 'John Doe',",
			},
		},
		"Renders custom user vars successfully": {
			userID:       "123456789",
			displayName:  "John Doe",
			emailAddress: "johndoe@example.com",
			opts: []IdentifyOption{
				IdentifyUserVars(map[string]interface{}{
					"foo":            "bar",
					"meaningOfLife":  42,
					"gravity":        9.8,
					"favoriteColors": []string{"red", "green", "blue"},
				}),
			},
			wantErr: false,
			doContains: []string{
				"FS.identify('123456789', {",
				"displayName: 'John Doe',",
				"email: 'johndoe@example.com',",
				"foo_str: 'bar'",
				"meaningOfLife_int: 42",
				"gravity_real: 9.8",
				"favoriteColors_strs: ['red','green','blue']",
			},
		},
	}

	for name, testCase := range testTable {
		t.Run(name, func(t *testing.T) {
			identify, err := IdentifySnippet(testCase.userID, testCase.displayName, testCase.emailAddress, testCase.opts...)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			for _, expect := range testCase.doContains {
				if !strings.Contains(string(identify), expect) {
					t.Logf("identify: %s", identify)
					t.Errorf("expected identify to contain %q", expect)
				}
			}

			for _, expect := range testCase.dontContains {
				if strings.Contains(string(identify), expect) {
					t.Logf("identify: %s", identify)
					t.Errorf("expected identify to not contain %q", expect)
				}
			}
		})
	}
}
