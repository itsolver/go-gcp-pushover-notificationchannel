package pushover_notificationchannel

const (
	templateMessage string = `
ProjectID: {{ .ProjectID }}
State: <b>{{ .State }}</b>
Summary: <font color="#0000ff">{{ .Summary }}</font>
Period: {{ .Started }}-->{{ .Ended }}
Metadata:
{{- range $Key, $Val := .SystemLabels}}
  {{ $Key }}: {{ $Val }}
{{- end}}
`
)
