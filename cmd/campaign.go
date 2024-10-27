package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCampaignCommand() *cobra.Command {
	campaignCmd := &cobra.Command{
		Use:   "campaign",
		Short: "Manage campaigns",
	}

	campaignCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Creates a new campaign",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for creating a campaign
			fmt.Println("Creating a new campaign...")
		},
	})

	campaignCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "Lists all campaigns",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for listing campaigns
			fmt.Println("Listing campaigns...")
		},
	})

	return campaignCmd
}
