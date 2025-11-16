package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_verifyMacCmd = &cobra.Command{
	Use:   "verify-mac",
	Short: "Verifies a Message Authentication Code (MAC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_verifyMacCmd).Standalone()

	paymentCryptographyData_verifyMacCmd.Flags().String("key-identifier", "", "The `keyARN` of the encryption key that Amazon Web Services Payment Cryptography uses to verify MAC data.")
	paymentCryptographyData_verifyMacCmd.Flags().String("mac", "", "The MAC being verified.")
	paymentCryptographyData_verifyMacCmd.Flags().String("mac-length", "", "The length of the MAC.")
	paymentCryptographyData_verifyMacCmd.Flags().String("message-data", "", "The data on for which MAC is under verification.")
	paymentCryptographyData_verifyMacCmd.Flags().String("verification-attributes", "", "The attributes and data values to use for MAC verification within Amazon Web Services Payment Cryptography.")
	paymentCryptographyData_verifyMacCmd.MarkFlagRequired("key-identifier")
	paymentCryptographyData_verifyMacCmd.MarkFlagRequired("mac")
	paymentCryptographyData_verifyMacCmd.MarkFlagRequired("message-data")
	paymentCryptographyData_verifyMacCmd.MarkFlagRequired("verification-attributes")
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_verifyMacCmd)
}
