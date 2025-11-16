package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_stopPipelineExecutionCmd = &cobra.Command{
	Use:   "stop-pipeline-execution",
	Short: "Stops the specified pipeline execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_stopPipelineExecutionCmd).Standalone()

	codepipeline_stopPipelineExecutionCmd.Flags().Bool("abandon", false, "Use this option to stop the pipeline execution by abandoning, rather than finishing, in-progress actions.")
	codepipeline_stopPipelineExecutionCmd.Flags().Bool("no-abandon", false, "Use this option to stop the pipeline execution by abandoning, rather than finishing, in-progress actions.")
	codepipeline_stopPipelineExecutionCmd.Flags().String("pipeline-execution-id", "", "The ID of the pipeline execution to be stopped in the current stage.")
	codepipeline_stopPipelineExecutionCmd.Flags().String("pipeline-name", "", "The name of the pipeline to stop.")
	codepipeline_stopPipelineExecutionCmd.Flags().String("reason", "", "Use this option to enter comments, such as the reason the pipeline was stopped.")
	codepipeline_stopPipelineExecutionCmd.Flag("no-abandon").Hidden = true
	codepipeline_stopPipelineExecutionCmd.MarkFlagRequired("pipeline-execution-id")
	codepipeline_stopPipelineExecutionCmd.MarkFlagRequired("pipeline-name")
	codepipelineCmd.AddCommand(codepipeline_stopPipelineExecutionCmd)
}
