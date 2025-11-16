package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getContextKeysForCustomPolicyCmd = &cobra.Command{
	Use:   "get-context-keys-for-custom-policy",
	Short: "Gets a list of all of the context keys referenced in the input policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getContextKeysForCustomPolicyCmd).Standalone()

	iam_getContextKeysForCustomPolicyCmd.Flags().String("policy-input-list", "", "A list of policies for which you want the list of context keys referenced in those policies.")
	iam_getContextKeysForCustomPolicyCmd.MarkFlagRequired("policy-input-list")
	iamCmd.AddCommand(iam_getContextKeysForCustomPolicyCmd)
}
