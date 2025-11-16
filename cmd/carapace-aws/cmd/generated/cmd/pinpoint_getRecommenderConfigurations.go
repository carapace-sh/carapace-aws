package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getRecommenderConfigurationsCmd = &cobra.Command{
	Use:   "get-recommender-configurations",
	Short: "Retrieves information about all the recommender model configurations that are associated with your Amazon Pinpoint account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getRecommenderConfigurationsCmd).Standalone()

	pinpoint_getRecommenderConfigurationsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getRecommenderConfigurationsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
	pinpointCmd.AddCommand(pinpoint_getRecommenderConfigurationsCmd)
}
