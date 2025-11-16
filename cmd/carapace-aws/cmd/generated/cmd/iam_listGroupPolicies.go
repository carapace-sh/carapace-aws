package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listGroupPoliciesCmd = &cobra.Command{
	Use:   "list-group-policies",
	Short: "Lists the names of the inline policies that are embedded in the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listGroupPoliciesCmd).Standalone()

	iam_listGroupPoliciesCmd.Flags().String("group-name", "", "The name of the group to list policies for.")
	iam_listGroupPoliciesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listGroupPoliciesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listGroupPoliciesCmd.MarkFlagRequired("group-name")
	iamCmd.AddCommand(iam_listGroupPoliciesCmd)
}
