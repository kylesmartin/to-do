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
			prompt := promptui.Prompt{
				Label: "What do you need to do?",
				// inputted string must not be empty
				Validate: func(s string) error {
					if s == "" {
						return fmt.Errorf("input must not be empty")
					}
					return nil
				},
			}
			task, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("error running prompt: %w", err)
			}

			// add new task to list of tasks
			list.Add(work.Task{
				Description: task,
			})
			logs.Info().Str("Task", task).Msg("Task added to to-do list")
			return nil
		},
		work.SaveList,
	)
}
