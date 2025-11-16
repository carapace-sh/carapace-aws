package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_updatePipelineCmd = &cobra.Command{
	Use:   "update-pipeline",
	Short: "Updates a specified pipeline with edits or changes to its structure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_updatePipelineCmd).Standalone()

	codepipeline_updatePipelineCmd.Flags().String("pipeline", "", "The name of the pipeline to be updated.")
	codepipeline_updatePipelineCmd.MarkFlagRequired("pipeline")
	codepipelineCmd.AddCommand(codepipeline_updatePipelineCmd)
}
