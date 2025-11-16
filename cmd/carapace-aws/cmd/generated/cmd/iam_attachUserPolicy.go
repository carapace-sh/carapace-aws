package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_attachUserPolicyCmd = &cobra.Command{
	Use:   "attach-user-policy",
	Short: "Attaches the specified managed policy to the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_attachUserPolicyCmd).Standalone()

	iam_attachUserPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy you want to attach.")
	iam_attachUserPolicyCmd.Flags().String("user-name", "", "The name (friendly name, not ARN) of the IAM user to attach the policy to.")
	iam_attachUserPolicyCmd.MarkFlagRequired("policy-arn")
	iam_attachUserPolicyCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_attachUserPolicyCmd)
}
