package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createModelCardCmd = &cobra.Command{
	Use:   "create-model-card",
	Short: "Creates an Amazon SageMaker Model Card.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createModelCardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createModelCardCmd).Standalone()

		sagemaker_createModelCardCmd.Flags().String("content", "", "The content of the model card.")
		sagemaker_createModelCardCmd.Flags().String("model-card-name", "", "The unique name of the model card.")
		sagemaker_createModelCardCmd.Flags().String("model-card-status", "", "The approval status of the model card within your organization.")
		sagemaker_createModelCardCmd.Flags().String("security-config", "", "An optional Key Management Service key to encrypt, decrypt, and re-encrypt model card content for regulated workloads with highly sensitive data.")
		sagemaker_createModelCardCmd.Flags().String("tags", "", "Key-value pairs used to manage metadata for model cards.")
		sagemaker_createModelCardCmd.MarkFlagRequired("content")
		sagemaker_createModelCardCmd.MarkFlagRequired("model-card-name")
		sagemaker_createModelCardCmd.MarkFlagRequired("model-card-status")
	})
	sagemakerCmd.AddCommand(sagemaker_createModelCardCmd)
}
