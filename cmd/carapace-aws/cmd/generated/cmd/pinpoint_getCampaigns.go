package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getCampaignsCmd = &cobra.Command{
	Use:   "get-campaigns",
	Short: "Retrieves information about the status, configuration, and other settings for all the campaigns that are associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getCampaignsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getCampaignsCmd).Standalone()

		pinpoint_getCampaignsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getCampaignsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getCampaignsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
		pinpoint_getCampaignsCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getCampaignsCmd)
}
