package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deletePolicyCmd = &cobra.Command{
	Use:   "delete-policy",
	Short: "Deletes the specified managed policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deletePolicyCmd).Standalone()

	iam_deletePolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the IAM policy you want to delete.")
	iam_deletePolicyCmd.MarkFlagRequired("policy-arn")
	iamCmd.AddCommand(iam_deletePolicyCmd)
}
