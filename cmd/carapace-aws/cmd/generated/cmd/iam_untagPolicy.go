package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_untagPolicyCmd = &cobra.Command{
	Use:   "untag-policy",
	Short: "Removes the specified tags from the customer managed policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_untagPolicyCmd).Standalone()

	iam_untagPolicyCmd.Flags().String("policy-arn", "", "The ARN of the IAM customer managed policy from which you want to remove tags.")
	iam_untagPolicyCmd.Flags().String("tag-keys", "", "A list of key names as a simple array of strings.")
	iam_untagPolicyCmd.MarkFlagRequired("policy-arn")
	iam_untagPolicyCmd.MarkFlagRequired("tag-keys")
	iamCmd.AddCommand(iam_untagPolicyCmd)
}
