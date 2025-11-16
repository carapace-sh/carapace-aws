package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_encryptDataCmd = &cobra.Command{
	Use:   "encrypt-data",
	Short: "Encrypts plaintext data to ciphertext using a symmetric (TDES, AES), asymmetric (RSA), or derived (DUKPT or EMV) encryption key scheme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_encryptDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptographyData_encryptDataCmd).Standalone()

		paymentCryptographyData_encryptDataCmd.Flags().String("encryption-attributes", "", "The encryption key type and attributes for plaintext encryption.")
		paymentCryptographyData_encryptDataCmd.Flags().String("key-identifier", "", "The `keyARN` of the encryption key that Amazon Web Services Payment Cryptography uses for plaintext encryption.")
		paymentCryptographyData_encryptDataCmd.Flags().String("plain-text", "", "The plaintext to be encrypted.")
		paymentCryptographyData_encryptDataCmd.Flags().String("wrapped-key", "", "The WrappedKeyBlock containing the encryption key for plaintext encryption.")
		paymentCryptographyData_encryptDataCmd.MarkFlagRequired("encryption-attributes")
		paymentCryptographyData_encryptDataCmd.MarkFlagRequired("key-identifier")
		paymentCryptographyData_encryptDataCmd.MarkFlagRequired("plain-text")
	})
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_encryptDataCmd)
}
