package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_updateCampaignSourceCmd = &cobra.Command{
	Use:   "update-campaign-source",
	Short: "Updates the campaign source with a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_updateCampaignSourceCmd).Standalone()

	connectcampaignsv2_updateCampaignSourceCmd.Flags().String("id", "", "")
	connectcampaignsv2_updateCampaignSourceCmd.Flags().String("source", "", "")
	connectcampaignsv2_updateCampaignSourceCmd.MarkFlagRequired("id")
	connectcampaignsv2_updateCampaignSourceCmd.MarkFlagRequired("source")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_updateCampaignSourceCmd)
}
