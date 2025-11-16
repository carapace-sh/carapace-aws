package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_putKeyPolicyCmd = &cobra.Command{
	Use:   "put-key-policy",
	Short: "Attaches a key policy to the specified KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_putKeyPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_putKeyPolicyCmd).Standalone()

		kms_putKeyPolicyCmd.Flags().String("bypass-policy-lockout-safety-check", "", "Skips (\"bypasses\") the key policy lockout safety check.")
		kms_putKeyPolicyCmd.Flags().String("key-id", "", "Sets the key policy on the specified KMS key.")
		kms_putKeyPolicyCmd.Flags().String("policy", "", "The key policy to attach to the KMS key.")
		kms_putKeyPolicyCmd.Flags().String("policy-name", "", "The name of the key policy.")
		kms_putKeyPolicyCmd.MarkFlagRequired("key-id")
		kms_putKeyPolicyCmd.MarkFlagRequired("policy")
	})
	kmsCmd.AddCommand(kms_putKeyPolicyCmd)
}
