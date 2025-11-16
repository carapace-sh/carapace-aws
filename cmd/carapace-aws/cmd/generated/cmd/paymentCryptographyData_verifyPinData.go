package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_verifyPinDataCmd = &cobra.Command{
	Use:   "verify-pin-data",
	Short: "Verifies pin-related data such as PIN and PIN Offset using algorithms including VISA PVV and IBM3624.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_verifyPinDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptographyData_verifyPinDataCmd).Standalone()

		paymentCryptographyData_verifyPinDataCmd.Flags().String("dukpt-attributes", "", "The attributes and values for the DUKPT encrypted PIN block data.")
		paymentCryptographyData_verifyPinDataCmd.Flags().String("encrypted-pin-block", "", "The encrypted PIN block data that Amazon Web Services Payment Cryptography verifies.")
		paymentCryptographyData_verifyPinDataCmd.Flags().String("encryption-key-identifier", "", "The `keyARN` of the encryption key under which the PIN block data is encrypted.")
		paymentCryptographyData_verifyPinDataCmd.Flags().String("encryption-wrapped-key", "", "")
		paymentCryptographyData_verifyPinDataCmd.Flags().String("pin-block-format", "", "The PIN encoding format for pin data generation as specified in ISO 9564. Amazon Web Services Payment Cryptography supports `ISO_Format_0` and `ISO_Format_3`.")
		paymentCryptographyData_verifyPinDataCmd.Flags().String("pin-data-length", "", "The length of PIN being verified.")
		paymentCryptographyData_verifyPinDataCmd.Flags().String("primary-account-number", "", "The Primary Account Number (PAN), a unique identifier for a payment credit or debit card that associates the card with a specific account holder.")
		paymentCryptographyData_verifyPinDataCmd.Flags().String("verification-attributes", "", "The attributes and values for PIN data verification.")
		paymentCryptographyData_verifyPinDataCmd.Flags().String("verification-key-identifier", "", "The `keyARN` of the PIN verification key.")
		paymentCryptographyData_verifyPinDataCmd.MarkFlagRequired("encrypted-pin-block")
		paymentCryptographyData_verifyPinDataCmd.MarkFlagRequired("encryption-key-identifier")
		paymentCryptographyData_verifyPinDataCmd.MarkFlagRequired("pin-block-format")
		paymentCryptographyData_verifyPinDataCmd.MarkFlagRequired("verification-attributes")
		paymentCryptographyData_verifyPinDataCmd.MarkFlagRequired("verification-key-identifier")
	})
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_verifyPinDataCmd)
}
