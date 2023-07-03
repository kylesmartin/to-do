package cmd

const (
	FILE_NAME     = "tasks.json"
	LIST_TEMPLATE = `
To-do: {{range .Tasks}}
  ● {{.Description}}
{{end}}
`
)
