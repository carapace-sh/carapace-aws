package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_setDefaultPolicyVersionCmd = &cobra.Command{
	Use:   "set-default-policy-version",
	Short: "Sets the specified version of the specified policy as the policy's default (operative) version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_setDefaultPolicyVersionCmd).Standalone()

	iam_setDefaultPolicyVersionCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy whose default version you want to set.")
	iam_setDefaultPolicyVersionCmd.Flags().String("version-id", "", "The version of the policy to set as the default (operative) version.")
	iam_setDefaultPolicyVersionCmd.MarkFlagRequired("policy-arn")
	iam_setDefaultPolicyVersionCmd.MarkFlagRequired("version-id")
	iamCmd.AddCommand(iam_setDefaultPolicyVersionCmd)
}
