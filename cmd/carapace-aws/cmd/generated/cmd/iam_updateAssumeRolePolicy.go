package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateAssumeRolePolicyCmd = &cobra.Command{
	Use:   "update-assume-role-policy",
	Short: "Updates the policy that grants an IAM entity permission to assume a role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateAssumeRolePolicyCmd).Standalone()

	iam_updateAssumeRolePolicyCmd.Flags().String("policy-document", "", "The policy that grants an entity permission to assume the role.")
	iam_updateAssumeRolePolicyCmd.Flags().String("role-name", "", "The name of the role to update with the new policy.")
	iam_updateAssumeRolePolicyCmd.MarkFlagRequired("policy-document")
	iam_updateAssumeRolePolicyCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_updateAssumeRolePolicyCmd)
}
