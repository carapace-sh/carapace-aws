package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_startCampaignCmd = &cobra.Command{
	Use:   "start-campaign",
	Short: "Starts a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_startCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_startCampaignCmd).Standalone()

		connectcampaignsv2_startCampaignCmd.Flags().String("id", "", "")
		connectcampaignsv2_startCampaignCmd.MarkFlagRequired("id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_startCampaignCmd)
}
