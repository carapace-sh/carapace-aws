package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createRecommenderConfigurationCmd = &cobra.Command{
	Use:   "create-recommender-configuration",
	Short: "Creates an Amazon Pinpoint configuration for a recommender model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createRecommenderConfigurationCmd).Standalone()

	pinpoint_createRecommenderConfigurationCmd.Flags().String("create-recommender-configuration", "", "")
	pinpoint_createRecommenderConfigurationCmd.MarkFlagRequired("create-recommender-configuration")
	pinpointCmd.AddCommand(pinpoint_createRecommenderConfigurationCmd)
}
