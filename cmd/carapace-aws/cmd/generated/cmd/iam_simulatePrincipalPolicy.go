package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_simulatePrincipalPolicyCmd = &cobra.Command{
	Use:   "simulate-principal-policy",
	Short: "Simulate how a set of IAM policies attached to an IAM entity works with a list of API operations and Amazon Web Services resources to determine the policies' effective permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_simulatePrincipalPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_simulatePrincipalPolicyCmd).Standalone()

		iam_simulatePrincipalPolicyCmd.Flags().String("action-names", "", "A list of names of API operations to evaluate in the simulation.")
		iam_simulatePrincipalPolicyCmd.Flags().String("caller-arn", "", "The ARN of the IAM user that you want to specify as the simulated caller of the API operations.")
		iam_simulatePrincipalPolicyCmd.Flags().String("context-entries", "", "A list of context keys and corresponding values for the simulation to use.")
		iam_simulatePrincipalPolicyCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_simulatePrincipalPolicyCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_simulatePrincipalPolicyCmd.Flags().String("permissions-boundary-policy-input-list", "", "The IAM permissions boundary policy to simulate.")
		iam_simulatePrincipalPolicyCmd.Flags().String("policy-input-list", "", "An optional list of additional policy documents to include in the simulation.")
		iam_simulatePrincipalPolicyCmd.Flags().String("policy-source-arn", "", "The Amazon Resource Name (ARN) of a user, group, or role whose policies you want to include in the simulation.")
		iam_simulatePrincipalPolicyCmd.Flags().String("resource-arns", "", "A list of ARNs of Amazon Web Services resources to include in the simulation.")
		iam_simulatePrincipalPolicyCmd.Flags().String("resource-handling-option", "", "Specifies the type of simulation to run.")
		iam_simulatePrincipalPolicyCmd.Flags().String("resource-owner", "", "An Amazon Web Services account ID that specifies the owner of any simulated resource that does not identify its owner in the resource ARN.")
		iam_simulatePrincipalPolicyCmd.Flags().String("resource-policy", "", "A resource-based policy to include in the simulation provided as a string.")
		iam_simulatePrincipalPolicyCmd.MarkFlagRequired("action-names")
		iam_simulatePrincipalPolicyCmd.MarkFlagRequired("policy-source-arn")
	})
	iamCmd.AddCommand(iam_simulatePrincipalPolicyCmd)
}
