package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_generateCardValidationDataCmd = &cobra.Command{
	Use:   "generate-card-validation-data",
	Short: "Generates card-related validation data using algorithms such as Card Verification Values (CVV/CVV2), Dynamic Card Verification Values (dCVV/dCVV2), or Card Security Codes (CSC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_generateCardValidationDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptographyData_generateCardValidationDataCmd).Standalone()

		paymentCryptographyData_generateCardValidationDataCmd.Flags().String("generation-attributes", "", "The algorithm for generating CVV or CSC values for the card within Amazon Web Services Payment Cryptography.")
		paymentCryptographyData_generateCardValidationDataCmd.Flags().String("key-identifier", "", "The `keyARN` of the CVK encryption key that Amazon Web Services Payment Cryptography uses to generate card data.")
		paymentCryptographyData_generateCardValidationDataCmd.Flags().String("primary-account-number", "", "The Primary Account Number (PAN), a unique identifier for a payment credit or debit card that associates the card with a specific account holder.")
		paymentCryptographyData_generateCardValidationDataCmd.Flags().String("validation-data-length", "", "The length of the CVV or CSC to be generated.")
		paymentCryptographyData_generateCardValidationDataCmd.MarkFlagRequired("generation-attributes")
		paymentCryptographyData_generateCardValidationDataCmd.MarkFlagRequired("key-identifier")
		paymentCryptographyData_generateCardValidationDataCmd.MarkFlagRequired("primary-account-number")
	})
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_generateCardValidationDataCmd)
}
