package user

import (
	"fmt"
	"mp-cli/config"

	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig()
			if err != nil {
				fmt.Printf("Error loading config: %v\n", err)
				return
			}
			fmt.Printf("Creating a new user using config: %+v\n", cfg)
			// Implement user creation logic here
		},
	}
}
