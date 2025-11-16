package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_deleteCampaignCmd = &cobra.Command{
	Use:   "delete-campaign",
	Short: "Deletes a campaign from the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_deleteCampaignCmd).Standalone()

	connectcampaignsv2_deleteCampaignCmd.Flags().String("id", "", "")
	connectcampaignsv2_deleteCampaignCmd.MarkFlagRequired("id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_deleteCampaignCmd)
}
