package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_deleteCampaignCommunicationTimeCmd = &cobra.Command{
	Use:   "delete-campaign-communication-time",
	Short: "Deletes the communication time config for a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_deleteCampaignCommunicationTimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_deleteCampaignCommunicationTimeCmd).Standalone()

		connectcampaignsv2_deleteCampaignCommunicationTimeCmd.Flags().String("config", "", "")
		connectcampaignsv2_deleteCampaignCommunicationTimeCmd.Flags().String("id", "", "")
		connectcampaignsv2_deleteCampaignCommunicationTimeCmd.MarkFlagRequired("config")
		connectcampaignsv2_deleteCampaignCommunicationTimeCmd.MarkFlagRequired("id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_deleteCampaignCommunicationTimeCmd)
}
