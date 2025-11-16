package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Retrieves information about the specified managed policy, including the policy's default version and the total number of IAM users, groups, and roles to which the policy is attached.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getPolicyCmd).Standalone()

		iam_getPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the managed policy that you want information about.")
		iam_getPolicyCmd.MarkFlagRequired("policy-arn")
	})
	iamCmd.AddCommand(iam_getPolicyCmd)
}
