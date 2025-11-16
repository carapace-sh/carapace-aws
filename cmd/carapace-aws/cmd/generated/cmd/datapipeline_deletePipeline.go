package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_deletePipelineCmd = &cobra.Command{
	Use:   "delete-pipeline",
	Short: "Deletes a pipeline, its pipeline definition, and its run history.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_deletePipelineCmd).Standalone()

	datapipeline_deletePipelineCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
	datapipeline_deletePipelineCmd.MarkFlagRequired("pipeline-id")
	datapipelineCmd.AddCommand(datapipeline_deletePipelineCmd)
}
