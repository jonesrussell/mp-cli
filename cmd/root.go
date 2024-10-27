/*
Copyright © 2024 Russell Jones <499552+jonesrussell@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "mp-cli",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Add this to ensure the app stops after command execution
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Command execution complete, shutting down...")
			os.Exit(0)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func RegisterRootCommand() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.AddCommand(NewUserCommand())
	rootCmd.AddCommand(NewCampaignCommand())
}
