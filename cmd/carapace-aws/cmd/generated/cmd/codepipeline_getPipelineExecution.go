package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_getPipelineExecutionCmd = &cobra.Command{
	Use:   "get-pipeline-execution",
	Short: "Returns information about an execution of a pipeline, including details about artifacts, the pipeline execution ID, and the name, version, and status of the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_getPipelineExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_getPipelineExecutionCmd).Standalone()

		codepipeline_getPipelineExecutionCmd.Flags().String("pipeline-execution-id", "", "The ID of the pipeline execution about which you want to get execution details.")
		codepipeline_getPipelineExecutionCmd.Flags().String("pipeline-name", "", "The name of the pipeline about which you want to get execution details.")
		codepipeline_getPipelineExecutionCmd.MarkFlagRequired("pipeline-execution-id")
		codepipeline_getPipelineExecutionCmd.MarkFlagRequired("pipeline-name")
	})
	codepipelineCmd.AddCommand(codepipeline_getPipelineExecutionCmd)
}
