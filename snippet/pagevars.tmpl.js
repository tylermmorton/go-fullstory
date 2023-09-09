FS ? FS.setVars('page', {
  'pageName': '{{ .PageName }}',
  {{ if .Vars }}{{ .Vars -}}{{ end }}
}) : function(){ throw new Error('FS is not defined. Did you forget to add your recording snippet?'); }();