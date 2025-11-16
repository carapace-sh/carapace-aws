package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_deleteScalingPolicyCmd = &cobra.Command{
	Use:   "delete-scaling-policy",
	Short: "Deletes the specified scaling policy for an Application Auto Scaling scalable target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_deleteScalingPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_deleteScalingPolicyCmd).Standalone()

		applicationAutoscaling_deleteScalingPolicyCmd.Flags().String("policy-name", "", "The name of the scaling policy.")
		applicationAutoscaling_deleteScalingPolicyCmd.Flags().String("resource-id", "", "The identifier of the resource associated with the scalable target.")
		applicationAutoscaling_deleteScalingPolicyCmd.Flags().String("scalable-dimension", "", "The scalable dimension.")
		applicationAutoscaling_deleteScalingPolicyCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
		applicationAutoscaling_deleteScalingPolicyCmd.MarkFlagRequired("policy-name")
		applicationAutoscaling_deleteScalingPolicyCmd.MarkFlagRequired("resource-id")
		applicationAutoscaling_deleteScalingPolicyCmd.MarkFlagRequired("scalable-dimension")
		applicationAutoscaling_deleteScalingPolicyCmd.MarkFlagRequired("service-namespace")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_deleteScalingPolicyCmd)
}
