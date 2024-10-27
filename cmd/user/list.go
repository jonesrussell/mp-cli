package user

import (
	"fmt"
	"mp-cli/config"

	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all users",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig()
			if err != nil {
				fmt.Printf("Error loading config: %v\n", err)
				return
			}
			fmt.Printf("Listing users using config: %+v\n", cfg)
			// Implement user listing logic here
		},
	}
}
