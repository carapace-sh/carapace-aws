package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_deleteCampaignCommunicationLimitsCmd = &cobra.Command{
	Use:   "delete-campaign-communication-limits",
	Short: "Deletes the communication limits config for a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_deleteCampaignCommunicationLimitsCmd).Standalone()

	connectcampaignsv2_deleteCampaignCommunicationLimitsCmd.Flags().String("config", "", "")
	connectcampaignsv2_deleteCampaignCommunicationLimitsCmd.Flags().String("id", "", "")
	connectcampaignsv2_deleteCampaignCommunicationLimitsCmd.MarkFlagRequired("config")
	connectcampaignsv2_deleteCampaignCommunicationLimitsCmd.MarkFlagRequired("id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_deleteCampaignCommunicationLimitsCmd)
}
