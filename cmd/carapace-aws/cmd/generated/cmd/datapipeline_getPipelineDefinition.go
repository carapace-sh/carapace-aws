package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_getPipelineDefinitionCmd = &cobra.Command{
	Use:   "get-pipeline-definition",
	Short: "Gets the definition of the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_getPipelineDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_getPipelineDefinitionCmd).Standalone()

		datapipeline_getPipelineDefinitionCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
		datapipeline_getPipelineDefinitionCmd.Flags().String("version", "", "The version of the pipeline definition to retrieve.")
		datapipeline_getPipelineDefinitionCmd.MarkFlagRequired("pipeline-id")
	})
	datapipelineCmd.AddCommand(datapipeline_getPipelineDefinitionCmd)
}
