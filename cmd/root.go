package cmd

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// the logger is used throughout cmd
var logs *zerolog.Logger

// Execute configures and runs the root command
func Execute() {
	rootCmd := configureRootCmd()
	if err := rootCmd.Execute(); err != nil {
		// we use code 1 to indicate that an error has forced exit
		os.Exit(1)
	}
}

func configureRootCmd() *cobra.Command {
	// configure logger
	logs = configureLogger()

	// create root command
	var rootCmd = cobra.Command{
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

	return &rootCmd
}

func configureLogger() *zerolog.Logger {
	// we use a human readable logger that outputs to console
	l := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).Level(zerolog.InfoLevel).With().Timestamp().Caller().Logger()

	return &l
}
