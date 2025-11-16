package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listPolicyVersionsCmd = &cobra.Command{
	Use:   "list-policy-versions",
	Short: "Lists information about the versions of the specified managed policy, including the version that is currently set as the policy's default version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listPolicyVersionsCmd).Standalone()

	iam_listPolicyVersionsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listPolicyVersionsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listPolicyVersionsCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy for which you want the versions.")
	iam_listPolicyVersionsCmd.MarkFlagRequired("policy-arn")
	iamCmd.AddCommand(iam_listPolicyVersionsCmd)
}
