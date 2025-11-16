package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_retryStageExecutionCmd = &cobra.Command{
	Use:   "retry-stage-execution",
	Short: "You can retry a stage that has failed without having to run a pipeline again from the beginning.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_retryStageExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_retryStageExecutionCmd).Standalone()

		codepipeline_retryStageExecutionCmd.Flags().String("pipeline-execution-id", "", "The ID of the pipeline execution in the failed stage to be retried.")
		codepipeline_retryStageExecutionCmd.Flags().String("pipeline-name", "", "The name of the pipeline that contains the failed stage.")
		codepipeline_retryStageExecutionCmd.Flags().String("retry-mode", "", "The scope of the retry attempt.")
		codepipeline_retryStageExecutionCmd.Flags().String("stage-name", "", "The name of the failed stage to be retried.")
		codepipeline_retryStageExecutionCmd.MarkFlagRequired("pipeline-execution-id")
		codepipeline_retryStageExecutionCmd.MarkFlagRequired("pipeline-name")
		codepipeline_retryStageExecutionCmd.MarkFlagRequired("retry-mode")
		codepipeline_retryStageExecutionCmd.MarkFlagRequired("stage-name")
	})
	codepipelineCmd.AddCommand(codepipeline_retryStageExecutionCmd)
}
