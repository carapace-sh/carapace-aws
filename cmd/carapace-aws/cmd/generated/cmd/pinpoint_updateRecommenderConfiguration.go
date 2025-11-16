package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateRecommenderConfigurationCmd = &cobra.Command{
	Use:   "update-recommender-configuration",
	Short: "Updates an Amazon Pinpoint configuration for a recommender model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateRecommenderConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateRecommenderConfigurationCmd).Standalone()

		pinpoint_updateRecommenderConfigurationCmd.Flags().String("recommender-id", "", "The unique identifier for the recommender model configuration.")
		pinpoint_updateRecommenderConfigurationCmd.Flags().String("update-recommender-configuration", "", "")
		pinpoint_updateRecommenderConfigurationCmd.MarkFlagRequired("recommender-id")
		pinpoint_updateRecommenderConfigurationCmd.MarkFlagRequired("update-recommender-configuration")
	})
	pinpointCmd.AddCommand(pinpoint_updateRecommenderConfigurationCmd)
}
