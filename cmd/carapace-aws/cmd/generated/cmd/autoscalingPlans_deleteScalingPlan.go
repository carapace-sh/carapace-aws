package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingPlans_deleteScalingPlanCmd = &cobra.Command{
	Use:   "delete-scaling-plan",
	Short: "Deletes the specified scaling plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingPlans_deleteScalingPlanCmd).Standalone()

	autoscalingPlans_deleteScalingPlanCmd.Flags().String("scaling-plan-name", "", "The name of the scaling plan.")
	autoscalingPlans_deleteScalingPlanCmd.Flags().String("scaling-plan-version", "", "The version number of the scaling plan.")
	autoscalingPlans_deleteScalingPlanCmd.MarkFlagRequired("scaling-plan-name")
	autoscalingPlans_deleteScalingPlanCmd.MarkFlagRequired("scaling-plan-version")
	autoscalingPlansCmd.AddCommand(autoscalingPlans_deleteScalingPlanCmd)
}
