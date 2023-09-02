package fullstory

import (
	"strings"
	"testing"
)

func Test_PagevarsSnippet(t *testing.T) {
	testCases := map[string]struct {
		pageName     string
		opts         []PageVarsOption
		doContains   []string
		dontContains []string
	}{
		"Renders the pagevars snippet successfully with the default options": {
			pageName: "Home",
			opts:     []PageVarsOption{},
			doContains: []string{
				`FS.setVars('page', {`,
				`'pageName': 'Home',`,
			},
		},
		"Renders custom page vars successfully": {
			pageName: "Home",
			opts: []PageVarsOption{
				PageVarsWithVar("foo", "bar"),
				PageVarsWithVar("meaningOfLife", 42),
				PageVarsWithVar("gravity", 9.8),
				PageVarsWithVar("favoriteColors", []string{"red", "green", "blue"}),
			},
			doContains: []string{
				`FS.setVars('page', {`,
				`'pageName': 'Home',`,
				"foo_str: \"bar\"",
				"meaningOfLife_int: 42",
				"gravity_real: 9.8",
				"favoriteColors_strs: [\"red\",\"green\",\"blue\"]",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			snippet, err := PageVarsSnippet(testCase.pageName, testCase.opts...)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			for _, expect := range testCase.doContains {
				if !strings.Contains(string(snippet.AsJS()), expect) {
					t.Logf("pagevars: %s", snippet)
					t.Errorf("expected pagevars to contain %q", expect)
				}
			}

			for _, expect := range testCase.dontContains {
				if strings.Contains(string(snippet.AsJS()), expect) {
					t.Logf("pagevars: %s", snippet)
					t.Errorf("expected pagevars to not contain %q", expect)
				}
			}
		})
	}
}
