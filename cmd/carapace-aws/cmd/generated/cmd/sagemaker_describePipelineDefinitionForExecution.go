package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describePipelineDefinitionForExecutionCmd = &cobra.Command{
	Use:   "describe-pipeline-definition-for-execution",
	Short: "Describes the details of an execution's pipeline definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describePipelineDefinitionForExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describePipelineDefinitionForExecutionCmd).Standalone()

		sagemaker_describePipelineDefinitionForExecutionCmd.Flags().String("pipeline-execution-arn", "", "The Amazon Resource Name (ARN) of the pipeline execution.")
		sagemaker_describePipelineDefinitionForExecutionCmd.MarkFlagRequired("pipeline-execution-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_describePipelineDefinitionForExecutionCmd)
}
