{{- /*gotype:github.com/tylermmorton/go-fullstory/snippet.OnReady*/ -}}
window['_fs_ready'] = function() {
    {{ range .Snippets }}{{ .AsJS }}{{ end }}
};