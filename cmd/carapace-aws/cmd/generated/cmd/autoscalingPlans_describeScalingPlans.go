package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingPlans_describeScalingPlansCmd = &cobra.Command{
	Use:   "describe-scaling-plans",
	Short: "Describes one or more of your scaling plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingPlans_describeScalingPlansCmd).Standalone()

	autoscalingPlans_describeScalingPlansCmd.Flags().String("application-sources", "", "The sources for the applications (up to 10).")
	autoscalingPlans_describeScalingPlansCmd.Flags().String("max-results", "", "The maximum number of scalable resources to return.")
	autoscalingPlans_describeScalingPlansCmd.Flags().String("next-token", "", "The token for the next set of results.")
	autoscalingPlans_describeScalingPlansCmd.Flags().String("scaling-plan-names", "", "The names of the scaling plans (up to 10).")
	autoscalingPlans_describeScalingPlansCmd.Flags().String("scaling-plan-version", "", "The version number of the scaling plan.")
	autoscalingPlansCmd.AddCommand(autoscalingPlans_describeScalingPlansCmd)
}
