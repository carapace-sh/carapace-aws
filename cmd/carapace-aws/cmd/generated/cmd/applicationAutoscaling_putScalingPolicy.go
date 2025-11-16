package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_putScalingPolicyCmd = &cobra.Command{
	Use:   "put-scaling-policy",
	Short: "Creates or updates a scaling policy for an Application Auto Scaling scalable target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_putScalingPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_putScalingPolicyCmd).Standalone()

		applicationAutoscaling_putScalingPolicyCmd.Flags().String("policy-name", "", "The name of the scaling policy.")
		applicationAutoscaling_putScalingPolicyCmd.Flags().String("policy-type", "", "The scaling policy type.")
		applicationAutoscaling_putScalingPolicyCmd.Flags().String("predictive-scaling-policy-configuration", "", "The configuration of the predictive scaling policy.")
		applicationAutoscaling_putScalingPolicyCmd.Flags().String("resource-id", "", "The identifier of the resource associated with the scaling policy.")
		applicationAutoscaling_putScalingPolicyCmd.Flags().String("scalable-dimension", "", "The scalable dimension.")
		applicationAutoscaling_putScalingPolicyCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
		applicationAutoscaling_putScalingPolicyCmd.Flags().String("step-scaling-policy-configuration", "", "A step scaling policy.")
		applicationAutoscaling_putScalingPolicyCmd.Flags().String("target-tracking-scaling-policy-configuration", "", "A target tracking scaling policy.")
		applicationAutoscaling_putScalingPolicyCmd.MarkFlagRequired("policy-name")
		applicationAutoscaling_putScalingPolicyCmd.MarkFlagRequired("resource-id")
		applicationAutoscaling_putScalingPolicyCmd.MarkFlagRequired("scalable-dimension")
		applicationAutoscaling_putScalingPolicyCmd.MarkFlagRequired("service-namespace")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_putScalingPolicyCmd)
}
