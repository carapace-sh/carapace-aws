package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateCampaignCmd = &cobra.Command{
	Use:   "update-campaign",
	Short: "Updates the configuration and other settings for a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateCampaignCmd).Standalone()

	pinpoint_updateCampaignCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_updateCampaignCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
	pinpoint_updateCampaignCmd.Flags().String("write-campaign-request", "", "")
	pinpoint_updateCampaignCmd.MarkFlagRequired("application-id")
	pinpoint_updateCampaignCmd.MarkFlagRequired("campaign-id")
	pinpoint_updateCampaignCmd.MarkFlagRequired("write-campaign-request")
	pinpointCmd.AddCommand(pinpoint_updateCampaignCmd)
}
