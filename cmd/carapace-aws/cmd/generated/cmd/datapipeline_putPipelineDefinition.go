package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_putPipelineDefinitionCmd = &cobra.Command{
	Use:   "put-pipeline-definition",
	Short: "Adds tasks, schedules, and preconditions to the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_putPipelineDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_putPipelineDefinitionCmd).Standalone()

		datapipeline_putPipelineDefinitionCmd.Flags().String("parameter-objects", "", "The parameter objects used with the pipeline.")
		datapipeline_putPipelineDefinitionCmd.Flags().String("parameter-values", "", "The parameter values used with the pipeline.")
		datapipeline_putPipelineDefinitionCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
		datapipeline_putPipelineDefinitionCmd.Flags().String("pipeline-objects", "", "The objects that define the pipeline.")
		datapipeline_putPipelineDefinitionCmd.MarkFlagRequired("pipeline-id")
		datapipeline_putPipelineDefinitionCmd.MarkFlagRequired("pipeline-objects")
	})
	datapipelineCmd.AddCommand(datapipeline_putPipelineDefinitionCmd)
}
