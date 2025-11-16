package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listAttachedRolePoliciesCmd = &cobra.Command{
	Use:   "list-attached-role-policies",
	Short: "Lists all managed policies that are attached to the specified IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listAttachedRolePoliciesCmd).Standalone()

	iam_listAttachedRolePoliciesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listAttachedRolePoliciesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listAttachedRolePoliciesCmd.Flags().String("path-prefix", "", "The path prefix for filtering the results.")
	iam_listAttachedRolePoliciesCmd.Flags().String("role-name", "", "The name (friendly name, not ARN) of the role to list attached policies for.")
	iam_listAttachedRolePoliciesCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_listAttachedRolePoliciesCmd)
}
