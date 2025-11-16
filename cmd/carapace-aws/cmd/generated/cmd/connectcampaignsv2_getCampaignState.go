package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_getCampaignStateCmd = &cobra.Command{
	Use:   "get-campaign-state",
	Short: "Get state of a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_getCampaignStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_getCampaignStateCmd).Standalone()

		connectcampaignsv2_getCampaignStateCmd.Flags().String("id", "", "")
		connectcampaignsv2_getCampaignStateCmd.MarkFlagRequired("id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_getCampaignStateCmd)
}
