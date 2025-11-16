package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_getCampaignStateBatchCmd = &cobra.Command{
	Use:   "get-campaign-state-batch",
	Short: "Get state of campaigns for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_getCampaignStateBatchCmd).Standalone()

	connectcampaignsv2_getCampaignStateBatchCmd.Flags().String("campaign-ids", "", "")
	connectcampaignsv2_getCampaignStateBatchCmd.MarkFlagRequired("campaign-ids")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_getCampaignStateBatchCmd)
}
