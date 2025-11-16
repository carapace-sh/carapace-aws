package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts ciphertext that was encrypted by a KMS key using any of the following operations:\n\n- [Encrypt]()\n- [GenerateDataKey]()\n- [GenerateDataKeyPair]()\n- [GenerateDataKeyWithoutPlaintext]()\n- [GenerateDataKeyPairWithoutPlaintext]()\n\nYou can use this operation to decrypt ciphertext that was encrypted under a symmetric encryption KMS key or an asymmetric encryption KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_decryptCmd).Standalone()

	kms_decryptCmd.Flags().String("ciphertext-blob", "", "Ciphertext to be decrypted.")
	kms_decryptCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_decryptCmd.Flags().String("encryption-algorithm", "", "Specifies the encryption algorithm that will be used to decrypt the ciphertext.")
	kms_decryptCmd.Flags().String("encryption-context", "", "Specifies the encryption context to use when decrypting the data.")
	kms_decryptCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_decryptCmd.Flags().String("key-id", "", "Specifies the KMS key that KMS uses to decrypt the ciphertext.")
	kms_decryptCmd.Flags().String("recipient", "", "A signed [attestation document](https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave-concepts.html#term-attestdoc) from an Amazon Web Services Nitro enclave or NitroTPM, and the encryption algorithm to use with the public key in the attestation document.")
	kms_decryptCmd.MarkFlagRequired("ciphertext-blob")
	kmsCmd.AddCommand(kms_decryptCmd)
}
