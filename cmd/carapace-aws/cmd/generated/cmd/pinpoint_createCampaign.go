package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createCampaignCmd = &cobra.Command{
	Use:   "create-campaign",
	Short: "Creates a new campaign for an application or updates the settings of an existing campaign for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_createCampaignCmd).Standalone()

		pinpoint_createCampaignCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_createCampaignCmd.Flags().String("write-campaign-request", "", "")
		pinpoint_createCampaignCmd.MarkFlagRequired("application-id")
		pinpoint_createCampaignCmd.MarkFlagRequired("write-campaign-request")
	})
	pinpointCmd.AddCommand(pinpoint_createCampaignCmd)
}
