package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteRecommenderConfigurationCmd = &cobra.Command{
	Use:   "delete-recommender-configuration",
	Short: "Deletes an Amazon Pinpoint configuration for a recommender model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteRecommenderConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteRecommenderConfigurationCmd).Standalone()

		pinpoint_deleteRecommenderConfigurationCmd.Flags().String("recommender-id", "", "The unique identifier for the recommender model configuration.")
		pinpoint_deleteRecommenderConfigurationCmd.MarkFlagRequired("recommender-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteRecommenderConfigurationCmd)
}
