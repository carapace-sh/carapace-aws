package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deletePipelineCmd = &cobra.Command{
	Use:   "delete-pipeline",
	Short: "Deletes a pipeline if there are no running instances of the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deletePipelineCmd).Standalone()

	sagemaker_deletePipelineCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
	sagemaker_deletePipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline to delete.")
	sagemaker_deletePipelineCmd.MarkFlagRequired("client-request-token")
	sagemaker_deletePipelineCmd.MarkFlagRequired("pipeline-name")
	sagemakerCmd.AddCommand(sagemaker_deletePipelineCmd)
}
