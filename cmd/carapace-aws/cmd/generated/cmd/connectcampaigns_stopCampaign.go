package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_stopCampaignCmd = &cobra.Command{
	Use:   "stop-campaign",
	Short: "Stops a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_stopCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_stopCampaignCmd).Standalone()

		connectcampaigns_stopCampaignCmd.Flags().String("id", "", "")
		connectcampaigns_stopCampaignCmd.MarkFlagRequired("id")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_stopCampaignCmd)
}
