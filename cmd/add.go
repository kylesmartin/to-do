package cmd

import (
	"fmt"

	"github.com/kylesmartin/to-do/pkg/work"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// AddCmd prompts the user for a task and adds it to the to-do list
func AddCmd() *cobra.Command {
	return NewSubCmd(
		"add",
		"Adds a task to the to-do list",
		work.LoadList,
		[]Validation{},
		func(list *work.List) error {
			// prompt user for new task
			// no validation needed because any string is valid
			prompt := promptui.Prompt{
				Label: "What do you need to do?",
			}
			task, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("error running prompt: %w", err)
			}

			// add new task to list of tasks
			list.Add(work.Task{
				Description: task,
			})
			return nil
		},
		work.SaveList,
	)
}
