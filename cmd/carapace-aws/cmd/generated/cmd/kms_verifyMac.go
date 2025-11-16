package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_verifyMacCmd = &cobra.Command{
	Use:   "verify-mac",
	Short: "Verifies the hash-based message authentication code (HMAC) for a specified message, HMAC KMS key, and MAC algorithm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_verifyMacCmd).Standalone()

	kms_verifyMacCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_verifyMacCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_verifyMacCmd.Flags().String("key-id", "", "The KMS key that will be used in the verification.")
	kms_verifyMacCmd.Flags().String("mac", "", "The HMAC to verify.")
	kms_verifyMacCmd.Flags().String("mac-algorithm", "", "The MAC algorithm that will be used in the verification.")
	kms_verifyMacCmd.Flags().String("message", "", "The message that will be used in the verification.")
	kms_verifyMacCmd.MarkFlagRequired("key-id")
	kms_verifyMacCmd.MarkFlagRequired("mac")
	kms_verifyMacCmd.MarkFlagRequired("mac-algorithm")
	kms_verifyMacCmd.MarkFlagRequired("message")
	kmsCmd.AddCommand(kms_verifyMacCmd)
}
