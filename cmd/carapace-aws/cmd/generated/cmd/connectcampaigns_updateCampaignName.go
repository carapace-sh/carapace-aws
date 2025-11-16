package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_updateCampaignNameCmd = &cobra.Command{
	Use:   "update-campaign-name",
	Short: "Updates the name of a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_updateCampaignNameCmd).Standalone()

	connectcampaigns_updateCampaignNameCmd.Flags().String("id", "", "")
	connectcampaigns_updateCampaignNameCmd.Flags().String("name", "", "")
	connectcampaigns_updateCampaignNameCmd.MarkFlagRequired("id")
	connectcampaigns_updateCampaignNameCmd.MarkFlagRequired("name")
	connectcampaignsCmd.AddCommand(connectcampaigns_updateCampaignNameCmd)
}
