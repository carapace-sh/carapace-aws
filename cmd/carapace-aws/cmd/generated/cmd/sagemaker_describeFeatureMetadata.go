package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeFeatureMetadataCmd = &cobra.Command{
	Use:   "describe-feature-metadata",
	Short: "Shows the metadata for a feature within a feature group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeFeatureMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeFeatureMetadataCmd).Standalone()

		sagemaker_describeFeatureMetadataCmd.Flags().String("feature-group-name", "", "The name or Amazon Resource Name (ARN) of the feature group containing the feature.")
		sagemaker_describeFeatureMetadataCmd.Flags().String("feature-name", "", "The name of the feature.")
		sagemaker_describeFeatureMetadataCmd.MarkFlagRequired("feature-group-name")
		sagemaker_describeFeatureMetadataCmd.MarkFlagRequired("feature-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeFeatureMetadataCmd)
}
