package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_reEncryptCmd = &cobra.Command{
	Use:   "re-encrypt",
	Short: "Decrypts ciphertext and then reencrypts it entirely within KMS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_reEncryptCmd).Standalone()

	kms_reEncryptCmd.Flags().String("ciphertext-blob", "", "Ciphertext of the data to reencrypt.")
	kms_reEncryptCmd.Flags().String("destination-encryption-algorithm", "", "Specifies the encryption algorithm that KMS will use to reecrypt the data after it has decrypted it.")
	kms_reEncryptCmd.Flags().String("destination-encryption-context", "", "Specifies that encryption context to use when the reencrypting the data.")
	kms_reEncryptCmd.Flags().String("destination-key-id", "", "A unique identifier for the KMS key that is used to reencrypt the data.")
	kms_reEncryptCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_reEncryptCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_reEncryptCmd.Flags().String("source-encryption-algorithm", "", "Specifies the encryption algorithm that KMS will use to decrypt the ciphertext before it is reencrypted.")
	kms_reEncryptCmd.Flags().String("source-encryption-context", "", "Specifies the encryption context to use to decrypt the ciphertext.")
	kms_reEncryptCmd.Flags().String("source-key-id", "", "Specifies the KMS key that KMS will use to decrypt the ciphertext before it is re-encrypted.")
	kms_reEncryptCmd.MarkFlagRequired("ciphertext-blob")
	kms_reEncryptCmd.MarkFlagRequired("destination-key-id")
	kmsCmd.AddCommand(kms_reEncryptCmd)
}
