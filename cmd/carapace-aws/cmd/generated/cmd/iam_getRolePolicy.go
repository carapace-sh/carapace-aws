package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getRolePolicyCmd = &cobra.Command{
	Use:   "get-role-policy",
	Short: "Retrieves the specified inline policy document that is embedded with the specified IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getRolePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getRolePolicyCmd).Standalone()

		iam_getRolePolicyCmd.Flags().String("policy-name", "", "The name of the policy document to get.")
		iam_getRolePolicyCmd.Flags().String("role-name", "", "The name of the role associated with the policy.")
		iam_getRolePolicyCmd.MarkFlagRequired("policy-name")
		iam_getRolePolicyCmd.MarkFlagRequired("role-name")
	})
	iamCmd.AddCommand(iam_getRolePolicyCmd)
}
