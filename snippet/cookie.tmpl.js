{{- /*gotype:github.com/tylermmorton/go-fullstory/snippet.cookie*/ -}}
FS ? document.cookie = `{{.Cookie.Name}}={{.Cookie.Value}};
  {{- if ne .Cookie.Path ""}} Path={{.Cookie.Path}};{{end -}}
  {{- if .Cookie.Secure}} Secure;{{end -}}
  {{- if .Cookie.HttpOnly }} HttpOnly;{{end -}}
`: throw new Error('FS is not defined. Did you forget to add your recording snippet?');