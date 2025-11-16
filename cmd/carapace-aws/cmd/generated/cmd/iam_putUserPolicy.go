package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_putUserPolicyCmd = &cobra.Command{
	Use:   "put-user-policy",
	Short: "Adds or updates an inline policy document that is embedded in the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_putUserPolicyCmd).Standalone()

	iam_putUserPolicyCmd.Flags().String("policy-document", "", "The policy document.")
	iam_putUserPolicyCmd.Flags().String("policy-name", "", "The name of the policy document.")
	iam_putUserPolicyCmd.Flags().String("user-name", "", "The name of the user to associate the policy with.")
	iam_putUserPolicyCmd.MarkFlagRequired("policy-document")
	iam_putUserPolicyCmd.MarkFlagRequired("policy-name")
	iam_putUserPolicyCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_putUserPolicyCmd)
}
