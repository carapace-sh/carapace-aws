package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_deleteCampaignCmd = &cobra.Command{
	Use:   "delete-campaign",
	Short: "Deletes a campaign from the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_deleteCampaignCmd).Standalone()

	connectcampaigns_deleteCampaignCmd.Flags().String("id", "", "")
	connectcampaigns_deleteCampaignCmd.MarkFlagRequired("id")
	connectcampaignsCmd.AddCommand(connectcampaigns_deleteCampaignCmd)
}
