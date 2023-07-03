package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/kylesmartin/to-do/pkg/work"
	"github.com/spf13/cobra"
)

var prettyList = template.Must(template.New("prettyList").Parse(SHOW_TEMPLATE))

// ShowCmd prints the to-do list
func ShowCmd() *cobra.Command {
	return NewSubCmd(
		"sho",
		"Prints the current to-do list",
		work.LoadList,
		[]func(l *work.List) error{IsNonEmpty},
		func(list *work.List) error {
			if err := prettyList.Execute(os.Stdout, list); err != nil {
				return fmt.Errorf("error displaying to-do list: %w", err)
			}

			return nil
		},
		work.SaveList,
	)
}
