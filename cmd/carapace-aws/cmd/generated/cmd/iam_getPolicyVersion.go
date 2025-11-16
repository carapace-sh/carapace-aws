package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getPolicyVersionCmd = &cobra.Command{
	Use:   "get-policy-version",
	Short: "Retrieves information about the specified version of the specified managed policy, including the policy document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getPolicyVersionCmd).Standalone()

	iam_getPolicyVersionCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the managed policy that you want information about.")
	iam_getPolicyVersionCmd.Flags().String("version-id", "", "Identifies the policy version to retrieve.")
	iam_getPolicyVersionCmd.MarkFlagRequired("policy-arn")
	iam_getPolicyVersionCmd.MarkFlagRequired("version-id")
	iamCmd.AddCommand(iam_getPolicyVersionCmd)
}
