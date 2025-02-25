package pushover_notificationchannel

const (
	templateMessage string = `
{{ if .RevisionName }}Revision: {{ .RevisionName }}{{ else }}Cloud Run Update{{ end }}

Status: {{ .State }}
Project: {{ .ProjectID }}`
)
