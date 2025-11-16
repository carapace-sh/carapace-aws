package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getGroupPolicyCmd = &cobra.Command{
	Use:   "get-group-policy",
	Short: "Retrieves the specified inline policy document that is embedded in the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getGroupPolicyCmd).Standalone()

	iam_getGroupPolicyCmd.Flags().String("group-name", "", "The name of the group the policy is associated with.")
	iam_getGroupPolicyCmd.Flags().String("policy-name", "", "The name of the policy document to get.")
	iam_getGroupPolicyCmd.MarkFlagRequired("group-name")
	iam_getGroupPolicyCmd.MarkFlagRequired("policy-name")
	iamCmd.AddCommand(iam_getGroupPolicyCmd)
}
