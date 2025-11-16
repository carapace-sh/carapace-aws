package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeFeatureGroupCmd = &cobra.Command{
	Use:   "describe-feature-group",
	Short: "Use this operation to describe a `FeatureGroup`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeFeatureGroupCmd).Standalone()

	sagemaker_describeFeatureGroupCmd.Flags().String("feature-group-name", "", "The name or Amazon Resource Name (ARN) of the `FeatureGroup` you want described.")
	sagemaker_describeFeatureGroupCmd.Flags().String("next-token", "", "A token to resume pagination of the list of `Features` (`FeatureDefinitions`).")
	sagemaker_describeFeatureGroupCmd.MarkFlagRequired("feature-group-name")
	sagemakerCmd.AddCommand(sagemaker_describeFeatureGroupCmd)
}
