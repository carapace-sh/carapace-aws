package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_updateCampaignCommunicationLimitsCmd = &cobra.Command{
	Use:   "update-campaign-communication-limits",
	Short: "Updates the communication limits config for a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_updateCampaignCommunicationLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_updateCampaignCommunicationLimitsCmd).Standalone()

		connectcampaignsv2_updateCampaignCommunicationLimitsCmd.Flags().String("communication-limits-override", "", "")
		connectcampaignsv2_updateCampaignCommunicationLimitsCmd.Flags().String("id", "", "")
		connectcampaignsv2_updateCampaignCommunicationLimitsCmd.MarkFlagRequired("communication-limits-override")
		connectcampaignsv2_updateCampaignCommunicationLimitsCmd.MarkFlagRequired("id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_updateCampaignCommunicationLimitsCmd)
}
