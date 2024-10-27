package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUserCommand() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
	}

	userCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Creates a new user",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for creating a user
			fmt.Println("Creating a new user...")
		},
	})

	userCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "Lists all users",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for listing users
			fmt.Println("Listing users...")
		},
	})

	return userCmd
}
