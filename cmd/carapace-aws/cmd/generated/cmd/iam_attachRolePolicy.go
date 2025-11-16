package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_attachRolePolicyCmd = &cobra.Command{
	Use:   "attach-role-policy",
	Short: "Attaches the specified managed policy to the specified IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_attachRolePolicyCmd).Standalone()

	iam_attachRolePolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy you want to attach.")
	iam_attachRolePolicyCmd.Flags().String("role-name", "", "The name (friendly name, not ARN) of the role to attach the policy to.")
	iam_attachRolePolicyCmd.MarkFlagRequired("policy-arn")
	iam_attachRolePolicyCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_attachRolePolicyCmd)
}
