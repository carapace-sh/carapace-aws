package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_activatePipelineCmd = &cobra.Command{
	Use:   "activate-pipeline",
	Short: "Validates the specified pipeline and starts processing pipeline tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_activatePipelineCmd).Standalone()

	datapipeline_activatePipelineCmd.Flags().String("parameter-values", "", "A list of parameter values to pass to the pipeline at activation.")
	datapipeline_activatePipelineCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
	datapipeline_activatePipelineCmd.Flags().String("start-timestamp", "", "The date and time to resume the pipeline.")
	datapipeline_activatePipelineCmd.MarkFlagRequired("pipeline-id")
	datapipelineCmd.AddCommand(datapipeline_activatePipelineCmd)
}
