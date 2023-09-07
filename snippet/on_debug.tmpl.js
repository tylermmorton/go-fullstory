{{- /*gotype:github.com/tylermmorton/go-fullstory/snippet.OnReady*/ -}}
window['_fs_debug'] = function() {
    {{ range .Snippets }}{{ .AsJS }}{{ end }}
};