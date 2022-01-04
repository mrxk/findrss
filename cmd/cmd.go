package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "findrss",
	Short: "findrss",
	RunE:  Serve,
}

func Execute() error {
	return rootCmd.Execute()
}
