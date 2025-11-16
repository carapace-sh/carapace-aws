package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_detachRolePolicyCmd = &cobra.Command{
	Use:   "detach-role-policy",
	Short: "Removes the specified managed policy from the specified role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_detachRolePolicyCmd).Standalone()

	iam_detachRolePolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy you want to detach.")
	iam_detachRolePolicyCmd.Flags().String("role-name", "", "The name (friendly name, not ARN) of the IAM role to detach the policy from.")
	iam_detachRolePolicyCmd.MarkFlagRequired("policy-arn")
	iam_detachRolePolicyCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_detachRolePolicyCmd)
}
