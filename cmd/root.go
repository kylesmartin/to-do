package cmd

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	logs          *zerolog.Logger
	consoleWriter = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
)

// Execute configures and runs the root command
func Execute() {
	rootCmd := configureRootCmd()
	if err := rootCmd.Execute(); err != nil {
		// TODO: extract proper error code based on execution
		os.Exit(1)
	}
}

func configureRootCmd() *cobra.Command {
	// initialize logger (TODO: extract logger to its own package)
	l := zerolog.New(consoleWriter).Level(zerolog.InfoLevel).With().Timestamp().Caller().Logger()
	logs = &l

	// create root command
	var rootCmd = &cobra.Command{
		Use:   "to-do",
		Short: "To-do manages a simple list of tasks",
		Long:  `A tool that enables users to manage a simple to-do list in a command line workflow`,
	}

	// add sub commands
	rootCmd.AddCommand(
		AddCmd(),
		ShowCmd(),
		CompleteCmd(),
	)

	return rootCmd
}
