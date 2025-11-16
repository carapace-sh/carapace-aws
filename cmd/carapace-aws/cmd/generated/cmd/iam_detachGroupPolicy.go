package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_detachGroupPolicyCmd = &cobra.Command{
	Use:   "detach-group-policy",
	Short: "Removes the specified managed policy from the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_detachGroupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_detachGroupPolicyCmd).Standalone()

		iam_detachGroupPolicyCmd.Flags().String("group-name", "", "The name (friendly name, not ARN) of the IAM group to detach the policy from.")
		iam_detachGroupPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy you want to detach.")
		iam_detachGroupPolicyCmd.MarkFlagRequired("group-name")
		iam_detachGroupPolicyCmd.MarkFlagRequired("policy-arn")
	})
	iamCmd.AddCommand(iam_detachGroupPolicyCmd)
}
