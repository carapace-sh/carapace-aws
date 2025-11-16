package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getUserPolicyCmd = &cobra.Command{
	Use:   "get-user-policy",
	Short: "Retrieves the specified inline policy document that is embedded in the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getUserPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getUserPolicyCmd).Standalone()

		iam_getUserPolicyCmd.Flags().String("policy-name", "", "The name of the policy document to get.")
		iam_getUserPolicyCmd.Flags().String("user-name", "", "The name of the user who the policy is associated with.")
		iam_getUserPolicyCmd.MarkFlagRequired("policy-name")
		iam_getUserPolicyCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_getUserPolicyCmd)
}
