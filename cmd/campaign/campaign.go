package campaign

import (
	"github.com/spf13/cobra"
)

func NewCampaignCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "campaign",
		Short: "Manage campaigns",
		Long:  `The campaign command allows you to manage campaigns in the system.`,
	}

	cmd.AddCommand(newCreateCommand())
	cmd.AddCommand(newListCommand())

	return cmd
}
