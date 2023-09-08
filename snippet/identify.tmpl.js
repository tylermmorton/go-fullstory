FS ? FS.identify('{{ .UserID }}', {
    email: '{{ .EmailAddress }}',
    displayName: '{{ .DisplayName }}',
    {{ if .UserVars }}{{ .UserVars }}{{ end }}
}) : function(){ throw new Error('FS is not defined. Did you forget to add your recording snippet?'); }();
