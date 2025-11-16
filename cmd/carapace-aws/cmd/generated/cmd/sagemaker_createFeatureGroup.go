package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createFeatureGroupCmd = &cobra.Command{
	Use:   "create-feature-group",
	Short: "Create a new `FeatureGroup`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createFeatureGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createFeatureGroupCmd).Standalone()

		sagemaker_createFeatureGroupCmd.Flags().String("description", "", "A free-form description of a `FeatureGroup`.")
		sagemaker_createFeatureGroupCmd.Flags().String("event-time-feature-name", "", "The name of the feature that stores the `EventTime` of a `Record` in a `FeatureGroup`.")
		sagemaker_createFeatureGroupCmd.Flags().String("feature-definitions", "", "A list of `Feature` names and types.")
		sagemaker_createFeatureGroupCmd.Flags().String("feature-group-name", "", "The name of the `FeatureGroup`.")
		sagemaker_createFeatureGroupCmd.Flags().String("offline-store-config", "", "Use this to configure an `OfflineFeatureStore`.")
		sagemaker_createFeatureGroupCmd.Flags().String("online-store-config", "", "You can turn the `OnlineStore` on or off by specifying `True` for the `EnableOnlineStore` flag in `OnlineStoreConfig`.")
		sagemaker_createFeatureGroupCmd.Flags().String("record-identifier-feature-name", "", "The name of the `Feature` whose value uniquely identifies a `Record` defined in the `FeatureStore`.")
		sagemaker_createFeatureGroupCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the `OfflineStore` if an `OfflineStoreConfig` is provided.")
		sagemaker_createFeatureGroupCmd.Flags().String("tags", "", "Tags used to identify `Features` in each `FeatureGroup`.")
		sagemaker_createFeatureGroupCmd.Flags().String("throughput-config", "", "")
		sagemaker_createFeatureGroupCmd.MarkFlagRequired("event-time-feature-name")
		sagemaker_createFeatureGroupCmd.MarkFlagRequired("feature-definitions")
		sagemaker_createFeatureGroupCmd.MarkFlagRequired("feature-group-name")
		sagemaker_createFeatureGroupCmd.MarkFlagRequired("record-identifier-feature-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createFeatureGroupCmd)
}
