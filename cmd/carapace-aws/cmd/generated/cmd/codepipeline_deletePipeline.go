package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_deletePipelineCmd = &cobra.Command{
	Use:   "delete-pipeline",
	Short: "Deletes the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_deletePipelineCmd).Standalone()

	codepipeline_deletePipelineCmd.Flags().String("name", "", "The name of the pipeline to be deleted.")
	codepipeline_deletePipelineCmd.MarkFlagRequired("name")
	codepipelineCmd.AddCommand(codepipeline_deletePipelineCmd)
}
