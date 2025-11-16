package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updatePipelineVersionCmd = &cobra.Command{
	Use:   "update-pipeline-version",
	Short: "Updates a pipeline version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updatePipelineVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updatePipelineVersionCmd).Standalone()

		sagemaker_updatePipelineVersionCmd.Flags().String("pipeline-arn", "", "The Amazon Resource Name (ARN) of the pipeline.")
		sagemaker_updatePipelineVersionCmd.Flags().String("pipeline-version-description", "", "The description of the pipeline version.")
		sagemaker_updatePipelineVersionCmd.Flags().String("pipeline-version-display-name", "", "The display name of the pipeline version.")
		sagemaker_updatePipelineVersionCmd.Flags().String("pipeline-version-id", "", "The pipeline version ID to update.")
		sagemaker_updatePipelineVersionCmd.MarkFlagRequired("pipeline-arn")
		sagemaker_updatePipelineVersionCmd.MarkFlagRequired("pipeline-version-id")
	})
	sagemakerCmd.AddCommand(sagemaker_updatePipelineVersionCmd)
}
