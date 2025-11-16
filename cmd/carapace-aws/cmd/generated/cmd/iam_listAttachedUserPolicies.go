package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listAttachedUserPoliciesCmd = &cobra.Command{
	Use:   "list-attached-user-policies",
	Short: "Lists all managed policies that are attached to the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listAttachedUserPoliciesCmd).Standalone()

	iam_listAttachedUserPoliciesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listAttachedUserPoliciesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listAttachedUserPoliciesCmd.Flags().String("path-prefix", "", "The path prefix for filtering the results.")
	iam_listAttachedUserPoliciesCmd.Flags().String("user-name", "", "The name (friendly name, not ARN) of the user to list attached policies for.")
	iam_listAttachedUserPoliciesCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_listAttachedUserPoliciesCmd)
}
