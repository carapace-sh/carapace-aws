package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_updateCampaignNameCmd = &cobra.Command{
	Use:   "update-campaign-name",
	Short: "Updates the name of a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_updateCampaignNameCmd).Standalone()

	connectcampaignsv2_updateCampaignNameCmd.Flags().String("id", "", "")
	connectcampaignsv2_updateCampaignNameCmd.Flags().String("name", "", "")
	connectcampaignsv2_updateCampaignNameCmd.MarkFlagRequired("id")
	connectcampaignsv2_updateCampaignNameCmd.MarkFlagRequired("name")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_updateCampaignNameCmd)
}
