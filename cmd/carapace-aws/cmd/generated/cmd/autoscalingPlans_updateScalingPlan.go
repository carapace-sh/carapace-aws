package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingPlans_updateScalingPlanCmd = &cobra.Command{
	Use:   "update-scaling-plan",
	Short: "Updates the specified scaling plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingPlans_updateScalingPlanCmd).Standalone()

	autoscalingPlans_updateScalingPlanCmd.Flags().String("application-source", "", "A CloudFormation stack or set of tags.")
	autoscalingPlans_updateScalingPlanCmd.Flags().String("scaling-instructions", "", "The scaling instructions.")
	autoscalingPlans_updateScalingPlanCmd.Flags().String("scaling-plan-name", "", "The name of the scaling plan.")
	autoscalingPlans_updateScalingPlanCmd.Flags().String("scaling-plan-version", "", "The version number of the scaling plan.")
	autoscalingPlans_updateScalingPlanCmd.MarkFlagRequired("scaling-plan-name")
	autoscalingPlans_updateScalingPlanCmd.MarkFlagRequired("scaling-plan-version")
	autoscalingPlansCmd.AddCommand(autoscalingPlans_updateScalingPlanCmd)
}
