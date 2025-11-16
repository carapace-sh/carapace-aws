package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_createPipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "Creates a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_createPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_createPipelineCmd).Standalone()

		codepipeline_createPipelineCmd.Flags().String("pipeline", "", "Represents the structure of actions and stages to be performed in the pipeline.")
		codepipeline_createPipelineCmd.Flags().String("tags", "", "The tags for the pipeline.")
		codepipeline_createPipelineCmd.MarkFlagRequired("pipeline")
	})
	codepipelineCmd.AddCommand(codepipeline_createPipelineCmd)
}
