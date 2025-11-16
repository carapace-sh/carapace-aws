package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeModelCardCmd = &cobra.Command{
	Use:   "describe-model-card",
	Short: "Describes the content, creation time, and security configuration of an Amazon SageMaker Model Card.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeModelCardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeModelCardCmd).Standalone()

		sagemaker_describeModelCardCmd.Flags().String("model-card-name", "", "The name or Amazon Resource Name (ARN) of the model card to describe.")
		sagemaker_describeModelCardCmd.Flags().String("model-card-version", "", "The version of the model card to describe.")
		sagemaker_describeModelCardCmd.MarkFlagRequired("model-card-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeModelCardCmd)
}
