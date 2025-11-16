package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_verifyCardValidationDataCmd = &cobra.Command{
	Use:   "verify-card-validation-data",
	Short: "Verifies card-related validation data using algorithms such as Card Verification Values (CVV/CVV2), Dynamic Card Verification Values (dCVV/dCVV2) and Card Security Codes (CSC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_verifyCardValidationDataCmd).Standalone()

	paymentCryptographyData_verifyCardValidationDataCmd.Flags().String("key-identifier", "", "The `keyARN` of the CVK encryption key that Amazon Web Services Payment Cryptography uses to verify card data.")
	paymentCryptographyData_verifyCardValidationDataCmd.Flags().String("primary-account-number", "", "The Primary Account Number (PAN), a unique identifier for a payment credit or debit card that associates the card with a specific account holder.")
	paymentCryptographyData_verifyCardValidationDataCmd.Flags().String("validation-data", "", "The CVV or CSC value for use for card data verification within Amazon Web Services Payment Cryptography.")
	paymentCryptographyData_verifyCardValidationDataCmd.Flags().String("verification-attributes", "", "The algorithm to use for verification of card data within Amazon Web Services Payment Cryptography.")
	paymentCryptographyData_verifyCardValidationDataCmd.MarkFlagRequired("key-identifier")
	paymentCryptographyData_verifyCardValidationDataCmd.MarkFlagRequired("primary-account-number")
	paymentCryptographyData_verifyCardValidationDataCmd.MarkFlagRequired("validation-data")
	paymentCryptographyData_verifyCardValidationDataCmd.MarkFlagRequired("verification-attributes")
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_verifyCardValidationDataCmd)
}
