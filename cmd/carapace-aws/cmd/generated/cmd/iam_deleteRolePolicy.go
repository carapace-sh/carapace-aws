package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteRolePolicyCmd = &cobra.Command{
	Use:   "delete-role-policy",
	Short: "Deletes the specified inline policy that is embedded in the specified IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteRolePolicyCmd).Standalone()

	iam_deleteRolePolicyCmd.Flags().String("policy-name", "", "The name of the inline policy to delete from the specified IAM role.")
	iam_deleteRolePolicyCmd.Flags().String("role-name", "", "The name (friendly name, not ARN) identifying the role that the policy is embedded in.")
	iam_deleteRolePolicyCmd.MarkFlagRequired("policy-name")
	iam_deleteRolePolicyCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_deleteRolePolicyCmd)
}
