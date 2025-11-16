package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getContextKeysForPrincipalPolicyCmd = &cobra.Command{
	Use:   "get-context-keys-for-principal-policy",
	Short: "Gets a list of all of the context keys referenced in all the IAM policies that are attached to the specified IAM entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getContextKeysForPrincipalPolicyCmd).Standalone()

	iam_getContextKeysForPrincipalPolicyCmd.Flags().String("policy-input-list", "", "An optional list of additional policies for which you want the list of context keys that are referenced.")
	iam_getContextKeysForPrincipalPolicyCmd.Flags().String("policy-source-arn", "", "The ARN of a user, group, or role whose policies contain the context keys that you want listed.")
	iam_getContextKeysForPrincipalPolicyCmd.MarkFlagRequired("policy-source-arn")
	iamCmd.AddCommand(iam_getContextKeysForPrincipalPolicyCmd)
}
