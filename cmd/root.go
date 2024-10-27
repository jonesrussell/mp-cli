/*
Copyright © 2024 Russell Jones <499552+jonesrussell@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"log"
	"mp-cli/config"

	"github.com/spf13/cobra"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "mp-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func RegisterRootCommand() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.AddCommand(NewHelloCommand())
}

func registerConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	return cfg
}
