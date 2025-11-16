package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingPlans_createScalingPlanCmd = &cobra.Command{
	Use:   "create-scaling-plan",
	Short: "Creates a scaling plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingPlans_createScalingPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscalingPlans_createScalingPlanCmd).Standalone()

		autoscalingPlans_createScalingPlanCmd.Flags().String("application-source", "", "A CloudFormation stack or set of tags.")
		autoscalingPlans_createScalingPlanCmd.Flags().String("scaling-instructions", "", "The scaling instructions.")
		autoscalingPlans_createScalingPlanCmd.Flags().String("scaling-plan-name", "", "The name of the scaling plan.")
		autoscalingPlans_createScalingPlanCmd.MarkFlagRequired("application-source")
		autoscalingPlans_createScalingPlanCmd.MarkFlagRequired("scaling-instructions")
		autoscalingPlans_createScalingPlanCmd.MarkFlagRequired("scaling-plan-name")
	})
	autoscalingPlansCmd.AddCommand(autoscalingPlans_createScalingPlanCmd)
}
