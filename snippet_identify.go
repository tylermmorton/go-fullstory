package fullstory

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	identifySnippetText = `
{{ if .Enabled -}}
FS.identify('{{ .UserID }}', {
  displayName: '{{ .DisplayName }}',
  email: '{{ .EmailAddress }}',
{{- if .UserVars }}
  {{ userVars .UserVars -}}
{{ end -}}
});
{{- end }}
`
	identifyTemplate = template.Must(template.New("fs-identify").Funcs(map[string]any{
		// https://help.fullstory.com/hc/en-us/articles/4446290296599-Setting-custom-API-properties
		"userVars": func(m map[string]interface{}) string {
			buf := strings.Builder{}
			for k, v := range m {
				switch v.(type) {
				case bool:
					buf.WriteString(fmt.Sprintf("%s_bool: %t,\n", k, v))
				case float32, float64:
					buf.WriteString(fmt.Sprintf("%s_real: %f,\n", k, v))
				case int, int32, int64:
					buf.WriteString(fmt.Sprintf("%s_int: %d,\n", k, v))
				case string:
					buf.WriteString(fmt.Sprintf("%s_str: %q,\n", k, v))
				case time.Time:
					buf.WriteString(fmt.Sprintf("%s_date: new Date(%q),\n", k, v.(time.Time).Format(time.RFC3339)))
				case []bool:
					bools := make([]string, len(v.([]bool)))
					for i, b := range v.([]bool) {
						bools[i] = strconv.FormatBool(b)
					}
					buf.WriteString(fmt.Sprintf("%s_bools: [%s],\n", k, strings.Join(bools, ",")))
				case []float32:
					reals := make([]string, len(v.([]float32)))
					for i, r := range v.([]float32) {
						reals[i] = fmt.Sprintf("%f", r)
					}
					buf.WriteString(fmt.Sprintf("%s_reals: [%s],\n", k, strings.Join(reals, ",")))
				case []float64:
					reals := make([]string, len(v.([]float64)))
					for i, r := range v.([]float64) {
						reals[i] = fmt.Sprintf("%f", r)
					}
					buf.WriteString(fmt.Sprintf("%s_reals: [%s],\n", k, strings.Join(reals, ",")))
				case []int:
					ints := make([]string, len(v.([]int)))
					for i, i2 := range v.([]int) {
						ints[i] = fmt.Sprintf("%d", i2)
					}
					buf.WriteString(fmt.Sprintf("%s_ints: [%s],\n", k, strings.Join(ints, ",")))
				case []int32:
					ints := make([]string, len(v.([]int32)))
					for i, i2 := range v.([]int32) {
						ints[i] = fmt.Sprintf("%d", i2)
					}
					buf.WriteString(fmt.Sprintf("%s_ints: [%s],\n", k, strings.Join(ints, ",")))
				case []int64:
					ints := make([]string, len(v.([]int64)))
					for i, i2 := range v.([]int64) {
						ints[i] = fmt.Sprintf("%d", i2)
					}
					buf.WriteString(fmt.Sprintf("%s_ints: [%s],\n", k, strings.Join(ints, ",")))
				case []string:
					strs := make([]string, len(v.([]string)))
					for i, s := range v.([]string) {
						strs[i] = fmt.Sprintf("%q", s)
					}
					buf.WriteString(fmt.Sprintf("%s_strs: [%s],\n", k, strings.Join(strs, ",")))
				case []time.Time:
					dates := make([]string, len(v.([]time.Time)))
					for i, d := range v.([]time.Time) {
						dates[i] = fmt.Sprintf("new Date(%q)", d.Format(time.RFC3339))
					}
					buf.WriteString(fmt.Sprintf("%s_dates: [%s]n", k, strings.Join(dates, ",")))
				}
			}
			return buf.String()
		},
	}).Parse(identifySnippetText))
)

type IdentifyOption func(*identify)

type identify struct {
	UserID       string
	DisplayName  string
	EmailAddress string

	Enabled  bool
	UserVars map[string]interface{}
}

func defaultIdentifyOptions(userID, displayName, emailAddress string) *identify {
	return &identify{
		Enabled:      true,
		UserID:       userID,
		DisplayName:  displayName,
		EmailAddress: emailAddress,
		UserVars:     nil,
	}
}

// IdentifyEnabled turns on/off the rendering of the entire <script> tag. It can be useful to
// control this via an environment variable.
//
//	fullstory.IdentifySnippet("123456789", "John Doe", "johndoe@example.com", fullstory.IdentifyEnabled(
//	  os.Getenv("FULLSTORY_ENABLED") == "true",
//	))
func IdentifyEnabled(enabled bool) IdentifyOption {
	return func(i *identify) {
		i.Enabled = enabled
	}
}

func IdentifyUserVars(userVars map[string]interface{}) IdentifyOption {
	return func(i *identify) {
		i.UserVars = userVars
	}
}

// IdentifySnippet renders a new FullStory identification recording within a <script> tag from a Go template.
// One can control how the identifier renders by supplying one or more IdentifyOptions
func IdentifySnippet(userID, displayName, emailAddress string, opts ...IdentifyOption) (Snippet, error) {
	dotCtx := defaultIdentifyOptions(userID, displayName, emailAddress)
	for _, opt := range opts {
		opt(dotCtx)
	}
	return createSnippetFromTemplate(identifyTemplate, dotCtx)
}

func MustIdentifySnippet(userID, displayName, emailAddress string, opts ...IdentifyOption) Snippet {
	html, err := IdentifySnippet(userID, displayName, emailAddress, opts...)
	if err != nil {
		panic(err)
	}
	return html
}
