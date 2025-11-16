package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateModelCardCmd = &cobra.Command{
	Use:   "update-model-card",
	Short: "Update an Amazon SageMaker Model Card.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateModelCardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateModelCardCmd).Standalone()

		sagemaker_updateModelCardCmd.Flags().String("content", "", "The updated model card content.")
		sagemaker_updateModelCardCmd.Flags().String("model-card-name", "", "The name or Amazon Resource Name (ARN) of the model card to update.")
		sagemaker_updateModelCardCmd.Flags().String("model-card-status", "", "The approval status of the model card within your organization.")
		sagemaker_updateModelCardCmd.MarkFlagRequired("model-card-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateModelCardCmd)
}
