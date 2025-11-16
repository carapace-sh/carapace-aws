package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_generatePinDataCmd = &cobra.Command{
	Use:   "generate-pin-data",
	Short: "Generates pin-related data such as PIN, PIN Verification Value (PVV), PIN Block, and PIN Offset during new card issuance or reissuance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_generatePinDataCmd).Standalone()

	paymentCryptographyData_generatePinDataCmd.Flags().String("encryption-key-identifier", "", "The `keyARN` of the PEK that Amazon Web Services Payment Cryptography uses to encrypt the PIN Block.")
	paymentCryptographyData_generatePinDataCmd.Flags().String("encryption-wrapped-key", "", "")
	paymentCryptographyData_generatePinDataCmd.Flags().String("generation-attributes", "", "The attributes and values to use for PIN, PVV, or PIN Offset generation.")
	paymentCryptographyData_generatePinDataCmd.Flags().String("generation-key-identifier", "", "The `keyARN` of the PEK that Amazon Web Services Payment Cryptography uses for pin data generation.")
	paymentCryptographyData_generatePinDataCmd.Flags().String("pin-block-format", "", "The PIN encoding format for pin data generation as specified in ISO 9564. Amazon Web Services Payment Cryptography supports `ISO_Format_0`, `ISO_Format_3` and `ISO_Format_4`.")
	paymentCryptographyData_generatePinDataCmd.Flags().String("pin-data-length", "", "The length of PIN under generation.")
	paymentCryptographyData_generatePinDataCmd.Flags().String("primary-account-number", "", "The Primary Account Number (PAN), a unique identifier for a payment credit or debit card that associates the card with a specific account holder.")
	paymentCryptographyData_generatePinDataCmd.MarkFlagRequired("encryption-key-identifier")
	paymentCryptographyData_generatePinDataCmd.MarkFlagRequired("generation-attributes")
	paymentCryptographyData_generatePinDataCmd.MarkFlagRequired("generation-key-identifier")
	paymentCryptographyData_generatePinDataCmd.MarkFlagRequired("pin-block-format")
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_generatePinDataCmd)
}
