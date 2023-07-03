package cmd

const (
	FILE_NAME     = "to-do.json"
	SHOW_TEMPLATE = `
To-do: {{range .Tasks}}
  ● {{.Description}}{{end}}

`
)
