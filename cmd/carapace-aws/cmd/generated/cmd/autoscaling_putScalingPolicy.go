package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_putScalingPolicyCmd = &cobra.Command{
	Use:   "put-scaling-policy",
	Short: "Creates or updates a scaling policy for an Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_putScalingPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_putScalingPolicyCmd).Standalone()

		autoscaling_putScalingPolicyCmd.Flags().String("adjustment-type", "", "Specifies how the scaling adjustment is interpreted (for example, an absolute number or a percentage).")
		autoscaling_putScalingPolicyCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_putScalingPolicyCmd.Flags().String("cooldown", "", "A cooldown period, in seconds, that applies to a specific simple scaling policy.")
		autoscaling_putScalingPolicyCmd.Flags().String("enabled", "", "Indicates whether the scaling policy is enabled or disabled.")
		autoscaling_putScalingPolicyCmd.Flags().String("estimated-instance-warmup", "", "*Not needed if the default instance warmup is defined for the group.*")
		autoscaling_putScalingPolicyCmd.Flags().String("metric-aggregation-type", "", "The aggregation type for the CloudWatch metrics.")
		autoscaling_putScalingPolicyCmd.Flags().String("min-adjustment-magnitude", "", "The minimum value to scale by when the adjustment type is `PercentChangeInCapacity`.")
		autoscaling_putScalingPolicyCmd.Flags().String("min-adjustment-step", "", "Available for backward compatibility.")
		autoscaling_putScalingPolicyCmd.Flags().String("policy-name", "", "The name of the policy.")
		autoscaling_putScalingPolicyCmd.Flags().String("policy-type", "", "One of the following policy types:")
		autoscaling_putScalingPolicyCmd.Flags().String("predictive-scaling-configuration", "", "A predictive scaling policy.")
		autoscaling_putScalingPolicyCmd.Flags().String("scaling-adjustment", "", "The amount by which to scale, based on the specified adjustment type.")
		autoscaling_putScalingPolicyCmd.Flags().String("step-adjustments", "", "A set of adjustments that enable you to scale based on the size of the alarm breach.")
		autoscaling_putScalingPolicyCmd.Flags().String("target-tracking-configuration", "", "A target tracking scaling policy.")
		autoscaling_putScalingPolicyCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_putScalingPolicyCmd.MarkFlagRequired("policy-name")
	})
	autoscalingCmd.AddCommand(autoscaling_putScalingPolicyCmd)
}
