package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteCampaignCmd = &cobra.Command{
	Use:   "delete-campaign",
	Short: "Deletes a campaign from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteCampaignCmd).Standalone()

	pinpoint_deleteCampaignCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_deleteCampaignCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
	pinpoint_deleteCampaignCmd.MarkFlagRequired("application-id")
	pinpoint_deleteCampaignCmd.MarkFlagRequired("campaign-id")
	pinpointCmd.AddCommand(pinpoint_deleteCampaignCmd)
}
