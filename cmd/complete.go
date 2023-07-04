package cmd

import (
	"fmt"

	"github.com/kylesmartin/to-do/pkg/work"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// CompleteCmd prompts the user to remove a task from the to-do list
func CompleteCmd() *cobra.Command {
	return NewSubCmd(
		"complete",
		"Removes a task from the to-do list",
		work.LoadList,
		[]Validation{IsNonEmpty},
		func(list *work.List) error {
			// get all descriptions
			var descriptions []string
			for _, task := range list.Tasks {
				descriptions = append(descriptions, task.Description)
			}

			// prompt user to select which task they want to remove
			prompt := promptui.Select{
				Label: "What have you completed?",
				Items: descriptions,
			}
			ind, task, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("error running selection: %w", err)
			}

			// remove task from list of tasks
			err = list.Remove(ind)
			if err == nil {
				logs.Info().Str("Task", task).Msg("Task removed from to-do list")
			}
			return err
		},
		work.SaveList,
	)
}
