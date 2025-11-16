package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_putGroupPolicyCmd = &cobra.Command{
	Use:   "put-group-policy",
	Short: "Adds or updates an inline policy document that is embedded in the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_putGroupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_putGroupPolicyCmd).Standalone()

		iam_putGroupPolicyCmd.Flags().String("group-name", "", "The name of the group to associate the policy with.")
		iam_putGroupPolicyCmd.Flags().String("policy-document", "", "The policy document.")
		iam_putGroupPolicyCmd.Flags().String("policy-name", "", "The name of the policy document.")
		iam_putGroupPolicyCmd.MarkFlagRequired("group-name")
		iam_putGroupPolicyCmd.MarkFlagRequired("policy-document")
		iam_putGroupPolicyCmd.MarkFlagRequired("policy-name")
	})
	iamCmd.AddCommand(iam_putGroupPolicyCmd)
}
