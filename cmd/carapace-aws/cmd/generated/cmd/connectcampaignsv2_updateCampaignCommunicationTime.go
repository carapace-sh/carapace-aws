package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_updateCampaignCommunicationTimeCmd = &cobra.Command{
	Use:   "update-campaign-communication-time",
	Short: "Updates the communication time config for a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_updateCampaignCommunicationTimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_updateCampaignCommunicationTimeCmd).Standalone()

		connectcampaignsv2_updateCampaignCommunicationTimeCmd.Flags().String("communication-time-config", "", "")
		connectcampaignsv2_updateCampaignCommunicationTimeCmd.Flags().String("id", "", "")
		connectcampaignsv2_updateCampaignCommunicationTimeCmd.MarkFlagRequired("communication-time-config")
		connectcampaignsv2_updateCampaignCommunicationTimeCmd.MarkFlagRequired("id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_updateCampaignCommunicationTimeCmd)
}
