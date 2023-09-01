package fullstory

import (
	"bytes"
	"html/template"
)

var (
	recordingSnippetText = /* .html */ `
{{ if .Enabled -}}
<script>
window['_fs_host'] = '{{ .Host }}';
window['_fs_script'] = '{{ .Script }}';
window['_fs_org'] = '{{ .OrgID }}';
window['_fs_namespace'] = '{{ .Namespace }}';
(function(m,n,e,t,l,o,g,y){
    if (e in m) {if(m.console && m.console.log) { m.console.log('FullStory namespace conflict. Please set window["_fs_namespace"].');} return;}
    g=m[e]=function(a,b,s){g.q?g.q.push([a,b,s]):g._api(a,b,s);};g.q=[];
    o=n.createElement(t);o.async=1;o.crossOrigin='anonymous';o.src='https://'+_fs_script;
    y=n.getElementsByTagName(t)[0];y.parentNode.insertBefore(o,y);
    g.identify=function(i,v,s){g(l,{uid:i},s);if(v)g(l,v,s)};g.setUserVars=function(v,s){g(l,v,s)};g.event=function(i,v,s){g('event',{n:i,p:v},s)};
    g.anonymize=function(){g.identify(!!0)};
    g.shutdown=function(){g("rec",!1)};g.restart=function(){g("rec",!0)};
    g.log = function(a,b){g("log",[a,b])};
    g.consent=function(a){g("consent",!arguments.length||a)};
    g.identifyAccount=function(i,v){o='account';v=v||{};v.acctId=i;g(o,v)};
    g.clearUserCookie=function(){};
    g.setVars=function(n, p){g('setVars',[n,p]);};
    g._w={};y='XMLHttpRequest';g._w[y]=m[y];y='fetch';g._w[y]=m[y];
    if(m[y])m[y]=function(){return g._w[y].apply(this,arguments)};
    g._v="1.3.0";
})(window,document,window['_fs_namespace'],'script','user');
</script>
{{- end }}
`
	recordingTemplate = template.Must(template.New("fs-recording").Parse(recordingSnippetText))
)

type RecordingOption func(*recording)

type recording struct {
	Enabled   bool
	Host      string
	Script    string
	OrgID     string
	Namespace string
}

func defaultRecordingOptions(orgID string) *recording {
	return &recording{
		Enabled:   true,
		Host:      "fullstory.com",
		Script:    "edge.fullstory.com/s/fs.js",
		OrgID:     orgID,
		Namespace: "FS",
	}
}

// RecordingEnabled turns on/off the rendering of the entire <script> tag. It can be useful to
// control this via an environment variable.
//
//	fullstory.RecordingSnippet(os.Getenv("FULLSTORY_ORG_ID"),
//	  fullstory.RecordingEnabled(os.Getenv("FULLSTORY_ENABLED") == "true"),
//	)
func RecordingEnabled(value bool) RecordingOption {
	return func(s *recording) {
		s.Enabled = value
	}
}

// RecordingHost sets the host of the FullStory recording recording.
func RecordingHost(value string) RecordingOption {
	return func(s *recording) {
		s.Host = value
	}
}

// RecordingScript sets the script URL of the FullStory recording recording.
func RecordingScript(value string) RecordingOption {
	return func(s *recording) {
		s.Script = value
	}
}

// RecordingNamespace sets the namespace of the FullStory recording recording.
func RecordingNamespace(value string) RecordingOption {
	return func(s *recording) {
		s.Namespace = value
	}
}

// RecordingSnippet renders a new FullStory recording recording within a <script> tag from a Go template.
// One can control how the recording renders by supplying one or more SnippetOptions
func RecordingSnippet(orgID string, opts ...RecordingOption) (template.HTML, error) {
	s := defaultRecordingOptions(orgID)
	for _, opt := range opts {
		opt(s)
	}
	buf := bytes.Buffer{}
	err := recordingTemplate.Execute(&buf, s)
	if err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}

func MustRecordingSnippet(orgID string, opts ...RecordingOption) template.HTML {
	snippet, err := RecordingSnippet(orgID, opts...)
	if err != nil {
		panic(err)
	}
	return snippet
}
