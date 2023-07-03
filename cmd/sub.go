package cmd

import (
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
	validations []func(*work.List) error,
	process func(*work.List) error,
	save func(string, *work.List) error,
) *cobra.Command {
	c := cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			// get home directory
			dir, err := os.UserHomeDir()
			if err != nil {
				logs.Fatal().Err(err).Msg("error getting home directory")
			}
			fileNameAndPath := path.Join(dir, FILE_NAME)

			// load list from file in home directory
			list, err := load(fileNameAndPath)
			if err != nil {
				logs.Fatal().Err(err).Msg("error loading to-do list")
			}

			// run validations
			for _, validation := range validations {
				err := validation(list)
				if err != nil {
					logs.Fatal().Err(err).Msg("validation failed")
				}
			}

			// perform some operation on the list
			err = process(list)
			if err != nil {
				logs.Fatal().Err(err).Msg("error processing to-do list")
			}

			// save the list after processing
			err = save(fileNameAndPath, list)
			if err != nil {
				logs.Fatal().Err(err).Msg("error saving to-do list")
			}
		},
	}

	return &c
}
