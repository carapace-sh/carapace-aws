package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_deactivatePipelineCmd = &cobra.Command{
	Use:   "deactivate-pipeline",
	Short: "Deactivates the specified running pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_deactivatePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_deactivatePipelineCmd).Standalone()

		datapipeline_deactivatePipelineCmd.Flags().String("cancel-active", "", "Indicates whether to cancel any running objects.")
		datapipeline_deactivatePipelineCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
		datapipeline_deactivatePipelineCmd.MarkFlagRequired("pipeline-id")
	})
	datapipelineCmd.AddCommand(datapipeline_deactivatePipelineCmd)
}
