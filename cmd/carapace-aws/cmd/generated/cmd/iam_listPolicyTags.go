package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listPolicyTagsCmd = &cobra.Command{
	Use:   "list-policy-tags",
	Short: "Lists the tags that are attached to the specified IAM customer managed policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listPolicyTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listPolicyTagsCmd).Standalone()

		iam_listPolicyTagsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listPolicyTagsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listPolicyTagsCmd.Flags().String("policy-arn", "", "The ARN of the IAM customer managed policy whose tags you want to see.")
		iam_listPolicyTagsCmd.MarkFlagRequired("policy-arn")
	})
	iamCmd.AddCommand(iam_listPolicyTagsCmd)
}
