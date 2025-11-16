package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateFeatureGroupCmd = &cobra.Command{
	Use:   "update-feature-group",
	Short: "Updates the feature group by either adding features or updating the online store configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateFeatureGroupCmd).Standalone()

	sagemaker_updateFeatureGroupCmd.Flags().String("feature-additions", "", "Updates the feature group.")
	sagemaker_updateFeatureGroupCmd.Flags().String("feature-group-name", "", "The name or Amazon Resource Name (ARN) of the feature group that you're updating.")
	sagemaker_updateFeatureGroupCmd.Flags().String("online-store-config", "", "Updates the feature group online store configuration.")
	sagemaker_updateFeatureGroupCmd.Flags().String("throughput-config", "", "")
	sagemaker_updateFeatureGroupCmd.MarkFlagRequired("feature-group-name")
	sagemakerCmd.AddCommand(sagemaker_updateFeatureGroupCmd)
}
