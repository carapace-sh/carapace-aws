package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getCampaignVersionCmd = &cobra.Command{
	Use:   "get-campaign-version",
	Short: "Retrieves information about the status, configuration, and other settings for a specific version of a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getCampaignVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getCampaignVersionCmd).Standalone()

		pinpoint_getCampaignVersionCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getCampaignVersionCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
		pinpoint_getCampaignVersionCmd.Flags().String("version", "", "The unique version number (Version property) for the campaign version.")
		pinpoint_getCampaignVersionCmd.MarkFlagRequired("application-id")
		pinpoint_getCampaignVersionCmd.MarkFlagRequired("campaign-id")
		pinpoint_getCampaignVersionCmd.MarkFlagRequired("version")
	})
	pinpointCmd.AddCommand(pinpoint_getCampaignVersionCmd)
}
