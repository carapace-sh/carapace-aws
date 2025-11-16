package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_pauseCampaignCmd = &cobra.Command{
	Use:   "pause-campaign",
	Short: "Pauses a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_pauseCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_pauseCampaignCmd).Standalone()

		connectcampaignsv2_pauseCampaignCmd.Flags().String("id", "", "")
		connectcampaignsv2_pauseCampaignCmd.MarkFlagRequired("id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_pauseCampaignCmd)
}
