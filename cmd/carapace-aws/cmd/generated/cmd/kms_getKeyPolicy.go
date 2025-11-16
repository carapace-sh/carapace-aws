package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_getKeyPolicyCmd = &cobra.Command{
	Use:   "get-key-policy",
	Short: "Gets a key policy attached to the specified KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_getKeyPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_getKeyPolicyCmd).Standalone()

		kms_getKeyPolicyCmd.Flags().String("key-id", "", "Gets the key policy for the specified KMS key.")
		kms_getKeyPolicyCmd.Flags().String("policy-name", "", "Specifies the name of the key policy.")
		kms_getKeyPolicyCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_getKeyPolicyCmd)
}
