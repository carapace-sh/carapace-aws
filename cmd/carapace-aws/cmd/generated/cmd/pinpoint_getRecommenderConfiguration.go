package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getRecommenderConfigurationCmd = &cobra.Command{
	Use:   "get-recommender-configuration",
	Short: "Retrieves information about an Amazon Pinpoint configuration for a recommender model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getRecommenderConfigurationCmd).Standalone()

	pinpoint_getRecommenderConfigurationCmd.Flags().String("recommender-id", "", "The unique identifier for the recommender model configuration.")
	pinpoint_getRecommenderConfigurationCmd.MarkFlagRequired("recommender-id")
	pinpointCmd.AddCommand(pinpoint_getRecommenderConfigurationCmd)
}
