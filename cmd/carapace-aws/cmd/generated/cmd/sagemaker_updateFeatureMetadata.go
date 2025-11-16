package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateFeatureMetadataCmd = &cobra.Command{
	Use:   "update-feature-metadata",
	Short: "Updates the description and parameters of the feature group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateFeatureMetadataCmd).Standalone()

	sagemaker_updateFeatureMetadataCmd.Flags().String("description", "", "A description that you can write to better describe the feature.")
	sagemaker_updateFeatureMetadataCmd.Flags().String("feature-group-name", "", "The name or Amazon Resource Name (ARN) of the feature group containing the feature that you're updating.")
	sagemaker_updateFeatureMetadataCmd.Flags().String("feature-name", "", "The name of the feature that you're updating.")
	sagemaker_updateFeatureMetadataCmd.Flags().String("parameter-additions", "", "A list of key-value pairs that you can add to better describe the feature.")
	sagemaker_updateFeatureMetadataCmd.Flags().String("parameter-removals", "", "A list of parameter keys that you can specify to remove parameters that describe your feature.")
	sagemaker_updateFeatureMetadataCmd.MarkFlagRequired("feature-group-name")
	sagemaker_updateFeatureMetadataCmd.MarkFlagRequired("feature-name")
	sagemakerCmd.AddCommand(sagemaker_updateFeatureMetadataCmd)
}
