package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_disableStageTransitionCmd = &cobra.Command{
	Use:   "disable-stage-transition",
	Short: "Prevents artifacts in a pipeline from transitioning to the next stage in the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_disableStageTransitionCmd).Standalone()

	codepipeline_disableStageTransitionCmd.Flags().String("pipeline-name", "", "The name of the pipeline in which you want to disable the flow of artifacts from one stage to another.")
	codepipeline_disableStageTransitionCmd.Flags().String("reason", "", "The reason given to the user that a stage is disabled, such as waiting for manual approval or manual tests.")
	codepipeline_disableStageTransitionCmd.Flags().String("stage-name", "", "The name of the stage where you want to disable the inbound or outbound transition of artifacts.")
	codepipeline_disableStageTransitionCmd.Flags().String("transition-type", "", "Specifies whether artifacts are prevented from transitioning into the stage and being processed by the actions in that stage (inbound), or prevented from transitioning from the stage after they have been processed by the actions in that stage (outbound).")
	codepipeline_disableStageTransitionCmd.MarkFlagRequired("pipeline-name")
	codepipeline_disableStageTransitionCmd.MarkFlagRequired("reason")
	codepipeline_disableStageTransitionCmd.MarkFlagRequired("stage-name")
	codepipeline_disableStageTransitionCmd.MarkFlagRequired("transition-type")
	codepipelineCmd.AddCommand(codepipeline_disableStageTransitionCmd)
}
