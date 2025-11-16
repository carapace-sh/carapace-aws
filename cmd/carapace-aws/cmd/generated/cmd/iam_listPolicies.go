package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listPoliciesCmd = &cobra.Command{
	Use:   "list-policies",
	Short: "Lists all the managed policies that are available in your Amazon Web Services account, including your own customer-defined managed policies and all Amazon Web Services managed policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listPoliciesCmd).Standalone()

		iam_listPoliciesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listPoliciesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listPoliciesCmd.Flags().String("only-attached", "", "A flag to filter the results to only the attached policies.")
		iam_listPoliciesCmd.Flags().String("path-prefix", "", "The path prefix for filtering the results.")
		iam_listPoliciesCmd.Flags().String("policy-usage-filter", "", "The policy usage method to use for filtering the results.")
		iam_listPoliciesCmd.Flags().String("scope", "", "The scope to use for filtering the results.")
	})
	iamCmd.AddCommand(iam_listPoliciesCmd)
}
