package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deletePolicyVersionCmd = &cobra.Command{
	Use:   "delete-policy-version",
	Short: "Deletes the specified version from the specified managed policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deletePolicyVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deletePolicyVersionCmd).Standalone()

		iam_deletePolicyVersionCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy from which you want to delete a version.")
		iam_deletePolicyVersionCmd.Flags().String("version-id", "", "The policy version to delete.")
		iam_deletePolicyVersionCmd.MarkFlagRequired("policy-arn")
		iam_deletePolicyVersionCmd.MarkFlagRequired("version-id")
	})
	iamCmd.AddCommand(iam_deletePolicyVersionCmd)
}
