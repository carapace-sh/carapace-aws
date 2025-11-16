package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_startPipelineExecutionCmd = &cobra.Command{
	Use:   "start-pipeline-execution",
	Short: "Starts a pipeline execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_startPipelineExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_startPipelineExecutionCmd).Standalone()

		sagemaker_startPipelineExecutionCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
		sagemaker_startPipelineExecutionCmd.Flags().String("parallelism-configuration", "", "This configuration, if specified, overrides the parallelism configuration of the parent pipeline for this specific run.")
		sagemaker_startPipelineExecutionCmd.Flags().String("pipeline-execution-description", "", "The description of the pipeline execution.")
		sagemaker_startPipelineExecutionCmd.Flags().String("pipeline-execution-display-name", "", "The display name of the pipeline execution.")
		sagemaker_startPipelineExecutionCmd.Flags().String("pipeline-name", "", "The name or Amazon Resource Name (ARN) of the pipeline.")
		sagemaker_startPipelineExecutionCmd.Flags().String("pipeline-parameters", "", "Contains a list of pipeline parameters.")
		sagemaker_startPipelineExecutionCmd.Flags().String("pipeline-version-id", "", "The ID of the pipeline version to start execution from.")
		sagemaker_startPipelineExecutionCmd.Flags().String("selective-execution-config", "", "The selective execution configuration applied to the pipeline run.")
		sagemaker_startPipelineExecutionCmd.MarkFlagRequired("client-request-token")
		sagemaker_startPipelineExecutionCmd.MarkFlagRequired("pipeline-name")
	})
	sagemakerCmd.AddCommand(sagemaker_startPipelineExecutionCmd)
}
