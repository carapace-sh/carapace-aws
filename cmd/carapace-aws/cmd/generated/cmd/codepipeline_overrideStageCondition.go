package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_overrideStageConditionCmd = &cobra.Command{
	Use:   "override-stage-condition",
	Short: "Used to override a stage condition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_overrideStageConditionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_overrideStageConditionCmd).Standalone()

		codepipeline_overrideStageConditionCmd.Flags().String("condition-type", "", "The type of condition to override for the stage, such as entry conditions, failure conditions, or success conditions.")
		codepipeline_overrideStageConditionCmd.Flags().String("pipeline-execution-id", "", "The ID of the pipeline execution for the override.")
		codepipeline_overrideStageConditionCmd.Flags().String("pipeline-name", "", "The name of the pipeline with the stage that will override the condition.")
		codepipeline_overrideStageConditionCmd.Flags().String("stage-name", "", "The name of the stage for the override.")
		codepipeline_overrideStageConditionCmd.MarkFlagRequired("condition-type")
		codepipeline_overrideStageConditionCmd.MarkFlagRequired("pipeline-execution-id")
		codepipeline_overrideStageConditionCmd.MarkFlagRequired("pipeline-name")
		codepipeline_overrideStageConditionCmd.MarkFlagRequired("stage-name")
	})
	codepipelineCmd.AddCommand(codepipeline_overrideStageConditionCmd)
}
