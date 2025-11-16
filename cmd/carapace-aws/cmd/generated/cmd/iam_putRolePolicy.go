package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_putRolePolicyCmd = &cobra.Command{
	Use:   "put-role-policy",
	Short: "Adds or updates an inline policy document that is embedded in the specified IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_putRolePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_putRolePolicyCmd).Standalone()

		iam_putRolePolicyCmd.Flags().String("policy-document", "", "The policy document.")
		iam_putRolePolicyCmd.Flags().String("policy-name", "", "The name of the policy document.")
		iam_putRolePolicyCmd.Flags().String("role-name", "", "The name of the role to associate the policy with.")
		iam_putRolePolicyCmd.MarkFlagRequired("policy-document")
		iam_putRolePolicyCmd.MarkFlagRequired("policy-name")
		iam_putRolePolicyCmd.MarkFlagRequired("role-name")
	})
	iamCmd.AddCommand(iam_putRolePolicyCmd)
}
