package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_startPipelineExecutionCmd = &cobra.Command{
	Use:   "start-pipeline-execution",
	Short: "Starts the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_startPipelineExecutionCmd).Standalone()

	codepipeline_startPipelineExecutionCmd.Flags().String("client-request-token", "", "The system-generated unique ID used to identify a unique execution request.")
	codepipeline_startPipelineExecutionCmd.Flags().String("name", "", "The name of the pipeline to start.")
	codepipeline_startPipelineExecutionCmd.Flags().String("source-revisions", "", "A list that allows you to specify, or override, the source revision for a pipeline execution that's being started.")
	codepipeline_startPipelineExecutionCmd.Flags().String("variables", "", "A list that overrides pipeline variables for a pipeline execution that's being started.")
	codepipeline_startPipelineExecutionCmd.MarkFlagRequired("name")
	codepipelineCmd.AddCommand(codepipeline_startPipelineExecutionCmd)
}
