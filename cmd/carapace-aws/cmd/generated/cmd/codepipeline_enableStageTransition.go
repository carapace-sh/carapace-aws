package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_enableStageTransitionCmd = &cobra.Command{
	Use:   "enable-stage-transition",
	Short: "Enables artifacts in a pipeline to transition to a stage in a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_enableStageTransitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_enableStageTransitionCmd).Standalone()

		codepipeline_enableStageTransitionCmd.Flags().String("pipeline-name", "", "The name of the pipeline in which you want to enable the flow of artifacts from one stage to another.")
		codepipeline_enableStageTransitionCmd.Flags().String("stage-name", "", "The name of the stage where you want to enable the transition of artifacts, either into the stage (inbound) or from that stage to the next stage (outbound).")
		codepipeline_enableStageTransitionCmd.Flags().String("transition-type", "", "Specifies whether artifacts are allowed to enter the stage and be processed by the actions in that stage (inbound) or whether already processed artifacts are allowed to transition to the next stage (outbound).")
		codepipeline_enableStageTransitionCmd.MarkFlagRequired("pipeline-name")
		codepipeline_enableStageTransitionCmd.MarkFlagRequired("stage-name")
		codepipeline_enableStageTransitionCmd.MarkFlagRequired("transition-type")
	})
	codepipelineCmd.AddCommand(codepipeline_enableStageTransitionCmd)
}
