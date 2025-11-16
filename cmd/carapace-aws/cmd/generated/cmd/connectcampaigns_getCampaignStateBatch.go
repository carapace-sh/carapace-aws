package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_getCampaignStateBatchCmd = &cobra.Command{
	Use:   "get-campaign-state-batch",
	Short: "Get state of campaigns for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_getCampaignStateBatchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_getCampaignStateBatchCmd).Standalone()

		connectcampaigns_getCampaignStateBatchCmd.Flags().String("campaign-ids", "", "")
		connectcampaigns_getCampaignStateBatchCmd.MarkFlagRequired("campaign-ids")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_getCampaignStateBatchCmd)
}
