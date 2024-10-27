package campaign

import (
	"fmt"
	"mp-cli/config"

	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new campaign",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig()
			if err != nil {
				fmt.Printf("Error loading config: %v\n", err)
				return
			}
			fmt.Printf("Creating a new campaign using config: %+v\n", cfg)
			// Implement campaign creation logic here
		},
	}
}
