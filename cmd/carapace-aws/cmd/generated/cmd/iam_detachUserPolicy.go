package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_detachUserPolicyCmd = &cobra.Command{
	Use:   "detach-user-policy",
	Short: "Removes the specified managed policy from the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_detachUserPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_detachUserPolicyCmd).Standalone()

		iam_detachUserPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy you want to detach.")
		iam_detachUserPolicyCmd.Flags().String("user-name", "", "The name (friendly name, not ARN) of the IAM user to detach the policy from.")
		iam_detachUserPolicyCmd.MarkFlagRequired("policy-arn")
		iam_detachUserPolicyCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_detachUserPolicyCmd)
}
