package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_listKeyPoliciesCmd = &cobra.Command{
	Use:   "list-key-policies",
	Short: "Gets the names of the key policies that are attached to a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_listKeyPoliciesCmd).Standalone()

	kms_listKeyPoliciesCmd.Flags().String("key-id", "", "Gets the names of key policies for the specified KMS key.")
	kms_listKeyPoliciesCmd.Flags().String("limit", "", "Use this parameter to specify the maximum number of items to return.")
	kms_listKeyPoliciesCmd.Flags().String("marker", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	kms_listKeyPoliciesCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_listKeyPoliciesCmd)
}
