package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_validatePipelineDefinitionCmd = &cobra.Command{
	Use:   "validate-pipeline-definition",
	Short: "Validates the specified pipeline definition to ensure that it is well formed and can be run without error.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_validatePipelineDefinitionCmd).Standalone()

	datapipeline_validatePipelineDefinitionCmd.Flags().String("parameter-objects", "", "The parameter objects used with the pipeline.")
	datapipeline_validatePipelineDefinitionCmd.Flags().String("parameter-values", "", "The parameter values used with the pipeline.")
	datapipeline_validatePipelineDefinitionCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
	datapipeline_validatePipelineDefinitionCmd.Flags().String("pipeline-objects", "", "The objects that define the pipeline changes to validate against the pipeline.")
	datapipeline_validatePipelineDefinitionCmd.MarkFlagRequired("pipeline-id")
	datapipeline_validatePipelineDefinitionCmd.MarkFlagRequired("pipeline-objects")
	datapipelineCmd.AddCommand(datapipeline_validatePipelineDefinitionCmd)
}
