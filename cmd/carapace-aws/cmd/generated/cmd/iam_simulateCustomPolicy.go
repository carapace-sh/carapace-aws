package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_simulateCustomPolicyCmd = &cobra.Command{
	Use:   "simulate-custom-policy",
	Short: "Simulate how a set of IAM policies and optionally a resource-based policy works with a list of API operations and Amazon Web Services resources to determine the policies' effective permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_simulateCustomPolicyCmd).Standalone()

	iam_simulateCustomPolicyCmd.Flags().String("action-names", "", "A list of names of API operations to evaluate in the simulation.")
	iam_simulateCustomPolicyCmd.Flags().String("caller-arn", "", "The ARN of the IAM user that you want to use as the simulated caller of the API operations.")
	iam_simulateCustomPolicyCmd.Flags().String("context-entries", "", "A list of context keys and corresponding values for the simulation to use.")
	iam_simulateCustomPolicyCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_simulateCustomPolicyCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_simulateCustomPolicyCmd.Flags().String("permissions-boundary-policy-input-list", "", "The IAM permissions boundary policy to simulate.")
	iam_simulateCustomPolicyCmd.Flags().String("policy-input-list", "", "A list of policy documents to include in the simulation.")
	iam_simulateCustomPolicyCmd.Flags().String("resource-arns", "", "A list of ARNs of Amazon Web Services resources to include in the simulation.")
	iam_simulateCustomPolicyCmd.Flags().String("resource-handling-option", "", "Specifies the type of simulation to run.")
	iam_simulateCustomPolicyCmd.Flags().String("resource-owner", "", "An ARN representing the Amazon Web Services account ID that specifies the owner of any simulated resource that does not identify its owner in the resource ARN.")
	iam_simulateCustomPolicyCmd.Flags().String("resource-policy", "", "A resource-based policy to include in the simulation provided as a string.")
	iam_simulateCustomPolicyCmd.MarkFlagRequired("action-names")
	iam_simulateCustomPolicyCmd.MarkFlagRequired("policy-input-list")
	iamCmd.AddCommand(iam_simulateCustomPolicyCmd)
}
