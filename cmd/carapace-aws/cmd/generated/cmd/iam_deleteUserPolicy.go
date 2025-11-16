package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteUserPolicyCmd = &cobra.Command{
	Use:   "delete-user-policy",
	Short: "Deletes the specified inline policy that is embedded in the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteUserPolicyCmd).Standalone()

	iam_deleteUserPolicyCmd.Flags().String("policy-name", "", "The name identifying the policy document to delete.")
	iam_deleteUserPolicyCmd.Flags().String("user-name", "", "The name (friendly name, not ARN) identifying the user that the policy is embedded in.")
	iam_deleteUserPolicyCmd.MarkFlagRequired("policy-name")
	iam_deleteUserPolicyCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_deleteUserPolicyCmd)
}
