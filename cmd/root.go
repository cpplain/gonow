package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gonow",
	Short: "A tool to sync ServiceNow records locally",
	Long: `Go Now is a Go-based CLI application to aid ServiceNow development.
The application enables developers to sync scripts locally for editing 
within any text editor and for version control via git.`,
}

func Execute() error {
	return rootCmd.Execute()
}
