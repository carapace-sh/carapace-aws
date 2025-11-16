package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_reEncryptDataCmd = &cobra.Command{
	Use:   "re-encrypt-data",
	Short: "Re-encrypt ciphertext using DUKPT or Symmetric data encryption keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_reEncryptDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptographyData_reEncryptDataCmd).Standalone()

		paymentCryptographyData_reEncryptDataCmd.Flags().String("cipher-text", "", "Ciphertext to be encrypted.")
		paymentCryptographyData_reEncryptDataCmd.Flags().String("incoming-encryption-attributes", "", "The attributes and values for incoming ciphertext.")
		paymentCryptographyData_reEncryptDataCmd.Flags().String("incoming-key-identifier", "", "The `keyARN` of the encryption key of incoming ciphertext data.")
		paymentCryptographyData_reEncryptDataCmd.Flags().String("incoming-wrapped-key", "", "The WrappedKeyBlock containing the encryption key of incoming ciphertext data.")
		paymentCryptographyData_reEncryptDataCmd.Flags().String("outgoing-encryption-attributes", "", "The attributes and values for outgoing ciphertext data after encryption by Amazon Web Services Payment Cryptography.")
		paymentCryptographyData_reEncryptDataCmd.Flags().String("outgoing-key-identifier", "", "The `keyARN` of the encryption key of outgoing ciphertext data after encryption by Amazon Web Services Payment Cryptography.")
		paymentCryptographyData_reEncryptDataCmd.Flags().String("outgoing-wrapped-key", "", "The WrappedKeyBlock containing the encryption key of outgoing ciphertext data after encryption by Amazon Web Services Payment Cryptography.")
		paymentCryptographyData_reEncryptDataCmd.MarkFlagRequired("cipher-text")
		paymentCryptographyData_reEncryptDataCmd.MarkFlagRequired("incoming-encryption-attributes")
		paymentCryptographyData_reEncryptDataCmd.MarkFlagRequired("incoming-key-identifier")
		paymentCryptographyData_reEncryptDataCmd.MarkFlagRequired("outgoing-encryption-attributes")
		paymentCryptographyData_reEncryptDataCmd.MarkFlagRequired("outgoing-key-identifier")
	})
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_reEncryptDataCmd)
}
