package cmd

const (
	FILE_NAME     = "tasks.json"
	SHOW_TEMPLATE = `
To-do: {{range .Tasks}}
  ● {{.Description}}{{end}}

`
)
