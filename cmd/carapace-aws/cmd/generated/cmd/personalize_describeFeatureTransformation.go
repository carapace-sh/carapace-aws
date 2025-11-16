package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeFeatureTransformationCmd = &cobra.Command{
	Use:   "describe-feature-transformation",
	Short: "Describes the given feature transformation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeFeatureTransformationCmd).Standalone()

	personalize_describeFeatureTransformationCmd.Flags().String("feature-transformation-arn", "", "The Amazon Resource Name (ARN) of the feature transformation to describe.")
	personalize_describeFeatureTransformationCmd.MarkFlagRequired("feature-transformation-arn")
	personalizeCmd.AddCommand(personalize_describeFeatureTransformationCmd)
}
