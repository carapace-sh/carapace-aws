package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_putScalingPolicyCmd = &cobra.Command{
	Use:   "put-scaling-policy",
	Short: "**This API works with the following fleet types:** EC2\n\nCreates or updates a scaling policy for a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_putScalingPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_putScalingPolicyCmd).Standalone()

		gamelift_putScalingPolicyCmd.Flags().String("comparison-operator", "", "Comparison operator to use when measuring the metric against the threshold value.")
		gamelift_putScalingPolicyCmd.Flags().String("evaluation-periods", "", "Length of time (in minutes) the metric must be at or beyond the threshold before a scaling event is triggered.")
		gamelift_putScalingPolicyCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to apply this policy to.")
		gamelift_putScalingPolicyCmd.Flags().String("metric-name", "", "Name of the Amazon GameLift Servers-defined metric that is used to trigger a scaling adjustment.")
		gamelift_putScalingPolicyCmd.Flags().String("name", "", "A descriptive label that is associated with a fleet's scaling policy.")
		gamelift_putScalingPolicyCmd.Flags().String("policy-type", "", "The type of scaling policy to create.")
		gamelift_putScalingPolicyCmd.Flags().String("scaling-adjustment", "", "Amount of adjustment to make, based on the scaling adjustment type.")
		gamelift_putScalingPolicyCmd.Flags().String("scaling-adjustment-type", "", "The type of adjustment to make to a fleet's instance count:")
		gamelift_putScalingPolicyCmd.Flags().String("target-configuration", "", "An object that contains settings for a target-based scaling policy.")
		gamelift_putScalingPolicyCmd.Flags().String("threshold", "", "Metric value used to trigger a scaling event.")
		gamelift_putScalingPolicyCmd.MarkFlagRequired("fleet-id")
		gamelift_putScalingPolicyCmd.MarkFlagRequired("metric-name")
		gamelift_putScalingPolicyCmd.MarkFlagRequired("name")
	})
	gameliftCmd.AddCommand(gamelift_putScalingPolicyCmd)
}
