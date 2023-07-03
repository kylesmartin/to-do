package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/kylesmartin/to-do/pkg/work"
	"github.com/spf13/cobra"
)

// NewSubCmd generates a subcommand that loads, processes, and saves a list
func NewSubCmd(
	use string,
	short string,
	load func(string) (*work.List, error),
	process func(*work.List) error,
	save func(string, *work.List) error,
) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			// get home directory
			dir, err := os.UserHomeDir()
			if err != nil {
				// TODO: use error types
				logs.Fatal().Err(fmt.Errorf("error getting home directory: %w", err))
			}
			fileNameAndPath := path.Join(dir, FILE_NAME)

			// load list from file in home directory
			list, err := load(fileNameAndPath)
			if err != nil {
				// TODO: use error types
				logs.Fatal().Err(fmt.Errorf("error loading to-do list: %w", err))
			}

			// perform some operation on the list
			err = process(list)
			if err != nil {
				// TODO: use error types
				logs.Fatal().Err(fmt.Errorf("error processing to-do list: %w", err))
			}

			// save the list after processing
			err = save(fileNameAndPath, list)
			if err != nil {
				// TODO: use error types
				logs.Fatal().Err(fmt.Errorf("error saving to-do list: %w", err))
			}
		},
	}
}
