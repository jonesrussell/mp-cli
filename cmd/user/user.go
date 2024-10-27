package user

import (
	"github.com/spf13/cobra"
)

func NewUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
		Long:  `The user command allows you to manage users in the system.`,
	}

	cmd.AddCommand(newCreateCommand())
	cmd.AddCommand(newListCommand())

	return cmd
}
