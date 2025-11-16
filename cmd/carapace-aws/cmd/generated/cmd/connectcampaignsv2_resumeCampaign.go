package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_resumeCampaignCmd = &cobra.Command{
	Use:   "resume-campaign",
	Short: "Stops a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_resumeCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_resumeCampaignCmd).Standalone()

		connectcampaignsv2_resumeCampaignCmd.Flags().String("id", "", "")
		connectcampaignsv2_resumeCampaignCmd.MarkFlagRequired("id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_resumeCampaignCmd)
}
