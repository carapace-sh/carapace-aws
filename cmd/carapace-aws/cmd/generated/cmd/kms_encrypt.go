package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts plaintext of up to 4,096 bytes using a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_encryptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_encryptCmd).Standalone()

		kms_encryptCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
		kms_encryptCmd.Flags().String("encryption-algorithm", "", "Specifies the encryption algorithm that KMS will use to encrypt the plaintext message.")
		kms_encryptCmd.Flags().String("encryption-context", "", "Specifies the encryption context that will be used to encrypt the data.")
		kms_encryptCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
		kms_encryptCmd.Flags().String("key-id", "", "Identifies the KMS key to use in the encryption operation.")
		kms_encryptCmd.Flags().String("plaintext", "", "Data to be encrypted.")
		kms_encryptCmd.MarkFlagRequired("key-id")
		kms_encryptCmd.MarkFlagRequired("plaintext")
	})
	kmsCmd.AddCommand(kms_encryptCmd)
}
