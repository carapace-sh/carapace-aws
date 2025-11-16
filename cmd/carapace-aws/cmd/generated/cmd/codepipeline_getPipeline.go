package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_getPipelineCmd = &cobra.Command{
	Use:   "get-pipeline",
	Short: "Returns the metadata, structure, stages, and actions of a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_getPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_getPipelineCmd).Standalone()

		codepipeline_getPipelineCmd.Flags().String("name", "", "The name of the pipeline for which you want to get information.")
		codepipeline_getPipelineCmd.Flags().String("version", "", "The version number of the pipeline.")
		codepipeline_getPipelineCmd.MarkFlagRequired("name")
	})
	codepipelineCmd.AddCommand(codepipeline_getPipelineCmd)
}
