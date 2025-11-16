package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getCampaignActivitiesCmd = &cobra.Command{
	Use:   "get-campaign-activities",
	Short: "Retrieves information about all the activities for a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getCampaignActivitiesCmd).Standalone()

	pinpoint_getCampaignActivitiesCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getCampaignActivitiesCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
	pinpoint_getCampaignActivitiesCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getCampaignActivitiesCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
	pinpoint_getCampaignActivitiesCmd.MarkFlagRequired("application-id")
	pinpoint_getCampaignActivitiesCmd.MarkFlagRequired("campaign-id")
	pinpointCmd.AddCommand(pinpoint_getCampaignActivitiesCmd)
}
