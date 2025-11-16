package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getCampaignCmd = &cobra.Command{
	Use:   "get-campaign",
	Short: "Retrieves information about the status, configuration, and other settings for a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getCampaignCmd).Standalone()

	pinpoint_getCampaignCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getCampaignCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
	pinpoint_getCampaignCmd.MarkFlagRequired("application-id")
	pinpoint_getCampaignCmd.MarkFlagRequired("campaign-id")
	pinpointCmd.AddCommand(pinpoint_getCampaignCmd)
}
