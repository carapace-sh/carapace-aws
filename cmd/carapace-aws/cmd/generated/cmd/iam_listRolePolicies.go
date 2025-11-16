package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listRolePoliciesCmd = &cobra.Command{
	Use:   "list-role-policies",
	Short: "Lists the names of the inline policies that are embedded in the specified IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listRolePoliciesCmd).Standalone()

	iam_listRolePoliciesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listRolePoliciesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listRolePoliciesCmd.Flags().String("role-name", "", "The name of the role to list policies for.")
	iam_listRolePoliciesCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_listRolePoliciesCmd)
}
