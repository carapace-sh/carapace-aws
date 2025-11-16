package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createPolicyVersionCmd = &cobra.Command{
	Use:   "create-policy-version",
	Short: "Creates a new version of the specified managed policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createPolicyVersionCmd).Standalone()

	iam_createPolicyVersionCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy to which you want to add a new version.")
	iam_createPolicyVersionCmd.Flags().String("policy-document", "", "The JSON policy document that you want to use as the content for this new version of the policy.")
	iam_createPolicyVersionCmd.Flags().String("set-as-default", "", "Specifies whether to set this version as the policy's default version.")
	iam_createPolicyVersionCmd.MarkFlagRequired("policy-arn")
	iam_createPolicyVersionCmd.MarkFlagRequired("policy-document")
	iamCmd.AddCommand(iam_createPolicyVersionCmd)
}
