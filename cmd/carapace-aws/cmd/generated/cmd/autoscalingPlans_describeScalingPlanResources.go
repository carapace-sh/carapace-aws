package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingPlans_describeScalingPlanResourcesCmd = &cobra.Command{
	Use:   "describe-scaling-plan-resources",
	Short: "Describes the scalable resources in the specified scaling plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingPlans_describeScalingPlanResourcesCmd).Standalone()

	autoscalingPlans_describeScalingPlanResourcesCmd.Flags().String("max-results", "", "The maximum number of scalable resources to return.")
	autoscalingPlans_describeScalingPlanResourcesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	autoscalingPlans_describeScalingPlanResourcesCmd.Flags().String("scaling-plan-name", "", "The name of the scaling plan.")
	autoscalingPlans_describeScalingPlanResourcesCmd.Flags().String("scaling-plan-version", "", "The version number of the scaling plan.")
	autoscalingPlans_describeScalingPlanResourcesCmd.MarkFlagRequired("scaling-plan-name")
	autoscalingPlans_describeScalingPlanResourcesCmd.MarkFlagRequired("scaling-plan-version")
	autoscalingPlansCmd.AddCommand(autoscalingPlans_describeScalingPlanResourcesCmd)
}
