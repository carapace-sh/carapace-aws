package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_describeScalingPoliciesCmd = &cobra.Command{
	Use:   "describe-scaling-policies",
	Short: "Describes the Application Auto Scaling scaling policies for the specified service namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_describeScalingPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_describeScalingPoliciesCmd).Standalone()

		applicationAutoscaling_describeScalingPoliciesCmd.Flags().String("max-results", "", "The maximum number of scalable targets.")
		applicationAutoscaling_describeScalingPoliciesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		applicationAutoscaling_describeScalingPoliciesCmd.Flags().String("policy-names", "", "The names of the scaling policies to describe.")
		applicationAutoscaling_describeScalingPoliciesCmd.Flags().String("resource-id", "", "The identifier of the resource associated with the scaling policy.")
		applicationAutoscaling_describeScalingPoliciesCmd.Flags().String("scalable-dimension", "", "The scalable dimension.")
		applicationAutoscaling_describeScalingPoliciesCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
		applicationAutoscaling_describeScalingPoliciesCmd.MarkFlagRequired("service-namespace")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_describeScalingPoliciesCmd)
}
