package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteGroupPolicyCmd = &cobra.Command{
	Use:   "delete-group-policy",
	Short: "Deletes the specified inline policy that is embedded in the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteGroupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deleteGroupPolicyCmd).Standalone()

		iam_deleteGroupPolicyCmd.Flags().String("group-name", "", "The name (friendly name, not ARN) identifying the group that the policy is embedded in.")
		iam_deleteGroupPolicyCmd.Flags().String("policy-name", "", "The name identifying the policy document to delete.")
		iam_deleteGroupPolicyCmd.MarkFlagRequired("group-name")
		iam_deleteGroupPolicyCmd.MarkFlagRequired("policy-name")
	})
	iamCmd.AddCommand(iam_deleteGroupPolicyCmd)
}
