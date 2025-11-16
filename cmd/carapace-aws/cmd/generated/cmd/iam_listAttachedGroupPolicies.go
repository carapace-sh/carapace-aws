package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listAttachedGroupPoliciesCmd = &cobra.Command{
	Use:   "list-attached-group-policies",
	Short: "Lists all managed policies that are attached to the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listAttachedGroupPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listAttachedGroupPoliciesCmd).Standalone()

		iam_listAttachedGroupPoliciesCmd.Flags().String("group-name", "", "The name (friendly name, not ARN) of the group to list attached policies for.")
		iam_listAttachedGroupPoliciesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listAttachedGroupPoliciesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listAttachedGroupPoliciesCmd.Flags().String("path-prefix", "", "The path prefix for filtering the results.")
		iam_listAttachedGroupPoliciesCmd.MarkFlagRequired("group-name")
	})
	iamCmd.AddCommand(iam_listAttachedGroupPoliciesCmd)
}
