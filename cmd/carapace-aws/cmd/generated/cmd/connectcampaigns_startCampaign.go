package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_startCampaignCmd = &cobra.Command{
	Use:   "start-campaign",
	Short: "Starts a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_startCampaignCmd).Standalone()

	connectcampaigns_startCampaignCmd.Flags().String("id", "", "")
	connectcampaigns_startCampaignCmd.MarkFlagRequired("id")
	connectcampaignsCmd.AddCommand(connectcampaigns_startCampaignCmd)
}
