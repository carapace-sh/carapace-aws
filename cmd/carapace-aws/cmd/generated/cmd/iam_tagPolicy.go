package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_tagPolicyCmd = &cobra.Command{
	Use:   "tag-policy",
	Short: "Adds one or more tags to an IAM customer managed policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_tagPolicyCmd).Standalone()

	iam_tagPolicyCmd.Flags().String("policy-arn", "", "The ARN of the IAM customer managed policy to which you want to add tags.")
	iam_tagPolicyCmd.Flags().String("tags", "", "The list of tags that you want to attach to the IAM customer managed policy.")
	iam_tagPolicyCmd.MarkFlagRequired("policy-arn")
	iam_tagPolicyCmd.MarkFlagRequired("tags")
	iamCmd.AddCommand(iam_tagPolicyCmd)
}
