package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_generateMacCmd = &cobra.Command{
	Use:   "generate-mac",
	Short: "Generates a hash-based message authentication code (HMAC) for a message using an HMAC KMS key and a MAC algorithm that the key supports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_generateMacCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_generateMacCmd).Standalone()

		kms_generateMacCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
		kms_generateMacCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
		kms_generateMacCmd.Flags().String("key-id", "", "The HMAC KMS key to use in the operation.")
		kms_generateMacCmd.Flags().String("mac-algorithm", "", "The MAC algorithm used in the operation.")
		kms_generateMacCmd.Flags().String("message", "", "The message to be hashed.")
		kms_generateMacCmd.MarkFlagRequired("key-id")
		kms_generateMacCmd.MarkFlagRequired("mac-algorithm")
		kms_generateMacCmd.MarkFlagRequired("message")
	})
	kmsCmd.AddCommand(kms_generateMacCmd)
}
