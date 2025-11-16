package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describePipelineExecutionCmd = &cobra.Command{
	Use:   "describe-pipeline-execution",
	Short: "Describes the details of a pipeline execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describePipelineExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describePipelineExecutionCmd).Standalone()

		sagemaker_describePipelineExecutionCmd.Flags().String("pipeline-execution-arn", "", "The Amazon Resource Name (ARN) of the pipeline execution.")
		sagemaker_describePipelineExecutionCmd.MarkFlagRequired("pipeline-execution-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_describePipelineExecutionCmd)
}
