{{ if .TYPE }}[{{ .TYPE }}]{{ end }}{{ if .ID }}[{{ .ID }}]{{ end }}{{ if .TITLE }}{{ .TITLE }}{{ end }}
{{ .summarize_prefix }}: {{ .summarize_title }}

{{ .summarize_message }}

{{ if .JIRA_URL }}{{ .JIRA_URL }}{{ end }}