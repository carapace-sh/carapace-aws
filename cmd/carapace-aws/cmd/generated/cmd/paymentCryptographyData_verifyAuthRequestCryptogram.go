package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_verifyAuthRequestCryptogramCmd = &cobra.Command{
	Use:   "verify-auth-request-cryptogram",
	Short: "Verifies Authorization Request Cryptogram (ARQC) for a EMV chip payment card authorization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_verifyAuthRequestCryptogramCmd).Standalone()

	paymentCryptographyData_verifyAuthRequestCryptogramCmd.Flags().String("auth-request-cryptogram", "", "The auth request cryptogram imported into Amazon Web Services Payment Cryptography for ARQC verification using a major encryption key and transaction data.")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.Flags().String("auth-response-attributes", "", "The attributes and values for auth request cryptogram verification.")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.Flags().String("key-identifier", "", "The `keyARN` of the major encryption key that Amazon Web Services Payment Cryptography uses for ARQC verification.")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.Flags().String("major-key-derivation-mode", "", "The method to use when deriving the major encryption key for ARQC verification within Amazon Web Services Payment Cryptography.")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.Flags().String("session-key-derivation-attributes", "", "The attributes and values to use for deriving a session key for ARQC verification within Amazon Web Services Payment Cryptography.")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.Flags().String("transaction-data", "", "The transaction data that Amazon Web Services Payment Cryptography uses for ARQC verification.")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.MarkFlagRequired("auth-request-cryptogram")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.MarkFlagRequired("key-identifier")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.MarkFlagRequired("major-key-derivation-mode")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.MarkFlagRequired("session-key-derivation-attributes")
	paymentCryptographyData_verifyAuthRequestCryptogramCmd.MarkFlagRequired("transaction-data")
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_verifyAuthRequestCryptogramCmd)
}
