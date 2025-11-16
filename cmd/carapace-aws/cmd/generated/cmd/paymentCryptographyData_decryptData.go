package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_decryptDataCmd = &cobra.Command{
	Use:   "decrypt-data",
	Short: "Decrypts ciphertext data to plaintext using a symmetric (TDES, AES), asymmetric (RSA), or derived (DUKPT or EMV) encryption key scheme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_decryptDataCmd).Standalone()

	paymentCryptographyData_decryptDataCmd.Flags().String("cipher-text", "", "The ciphertext to decrypt.")
	paymentCryptographyData_decryptDataCmd.Flags().String("decryption-attributes", "", "The encryption key type and attributes for ciphertext decryption.")
	paymentCryptographyData_decryptDataCmd.Flags().String("key-identifier", "", "The `keyARN` of the encryption key that Amazon Web Services Payment Cryptography uses for ciphertext decryption.")
	paymentCryptographyData_decryptDataCmd.Flags().String("wrapped-key", "", "The WrappedKeyBlock containing the encryption key for ciphertext decryption.")
	paymentCryptographyData_decryptDataCmd.MarkFlagRequired("cipher-text")
	paymentCryptographyData_decryptDataCmd.MarkFlagRequired("decryption-attributes")
	paymentCryptographyData_decryptDataCmd.MarkFlagRequired("key-identifier")
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_decryptDataCmd)
}
