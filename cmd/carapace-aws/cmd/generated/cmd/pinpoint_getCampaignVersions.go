package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getCampaignVersionsCmd = &cobra.Command{
	Use:   "get-campaign-versions",
	Short: "Retrieves information about the status, configuration, and other settings for all versions of a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getCampaignVersionsCmd).Standalone()

	pinpoint_getCampaignVersionsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getCampaignVersionsCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
	pinpoint_getCampaignVersionsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getCampaignVersionsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
	pinpoint_getCampaignVersionsCmd.MarkFlagRequired("application-id")
	pinpoint_getCampaignVersionsCmd.MarkFlagRequired("campaign-id")
	pinpointCmd.AddCommand(pinpoint_getCampaignVersionsCmd)
}
