package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updatePipelineExecutionCmd = &cobra.Command{
	Use:   "update-pipeline-execution",
	Short: "Updates a pipeline execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updatePipelineExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updatePipelineExecutionCmd).Standalone()

		sagemaker_updatePipelineExecutionCmd.Flags().String("parallelism-configuration", "", "This configuration, if specified, overrides the parallelism configuration of the parent pipeline for this specific run.")
		sagemaker_updatePipelineExecutionCmd.Flags().String("pipeline-execution-arn", "", "The Amazon Resource Name (ARN) of the pipeline execution.")
		sagemaker_updatePipelineExecutionCmd.Flags().String("pipeline-execution-description", "", "The description of the pipeline execution.")
		sagemaker_updatePipelineExecutionCmd.Flags().String("pipeline-execution-display-name", "", "The display name of the pipeline execution.")
		sagemaker_updatePipelineExecutionCmd.MarkFlagRequired("pipeline-execution-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_updatePipelineExecutionCmd)
}
