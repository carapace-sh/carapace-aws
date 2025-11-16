package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listEntitiesForPolicyCmd = &cobra.Command{
	Use:   "list-entities-for-policy",
	Short: "Lists all IAM users, groups, and roles that the specified managed policy is attached to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listEntitiesForPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listEntitiesForPolicyCmd).Standalone()

		iam_listEntitiesForPolicyCmd.Flags().String("entity-filter", "", "The entity type to use for filtering the results.")
		iam_listEntitiesForPolicyCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listEntitiesForPolicyCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listEntitiesForPolicyCmd.Flags().String("path-prefix", "", "The path prefix for filtering the results.")
		iam_listEntitiesForPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy for which you want the versions.")
		iam_listEntitiesForPolicyCmd.Flags().String("policy-usage-filter", "", "The policy usage method to use for filtering the results.")
		iam_listEntitiesForPolicyCmd.MarkFlagRequired("policy-arn")
	})
	iamCmd.AddCommand(iam_listEntitiesForPolicyCmd)
}
