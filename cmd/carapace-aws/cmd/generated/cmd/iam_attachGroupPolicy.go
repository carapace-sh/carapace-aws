package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_attachGroupPolicyCmd = &cobra.Command{
	Use:   "attach-group-policy",
	Short: "Attaches the specified managed policy to the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_attachGroupPolicyCmd).Standalone()

	iam_attachGroupPolicyCmd.Flags().String("group-name", "", "The name (friendly name, not ARN) of the group to attach the policy to.")
	iam_attachGroupPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy you want to attach.")
	iam_attachGroupPolicyCmd.MarkFlagRequired("group-name")
	iam_attachGroupPolicyCmd.MarkFlagRequired("policy-arn")
	iamCmd.AddCommand(iam_attachGroupPolicyCmd)
}
