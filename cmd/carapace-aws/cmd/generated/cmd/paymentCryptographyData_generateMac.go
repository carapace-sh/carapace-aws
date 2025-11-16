package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_generateMacCmd = &cobra.Command{
	Use:   "generate-mac",
	Short: "Generates a Message Authentication Code (MAC) cryptogram within Amazon Web Services Payment Cryptography.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_generateMacCmd).Standalone()

	paymentCryptographyData_generateMacCmd.Flags().String("generation-attributes", "", "The attributes and data values to use for MAC generation within Amazon Web Services Payment Cryptography.")
	paymentCryptographyData_generateMacCmd.Flags().String("key-identifier", "", "The `keyARN` of the MAC generation encryption key.")
	paymentCryptographyData_generateMacCmd.Flags().String("mac-length", "", "The length of a MAC under generation.")
	paymentCryptographyData_generateMacCmd.Flags().String("message-data", "", "The data for which a MAC is under generation.")
	paymentCryptographyData_generateMacCmd.MarkFlagRequired("generation-attributes")
	paymentCryptographyData_generateMacCmd.MarkFlagRequired("key-identifier")
	paymentCryptographyData_generateMacCmd.MarkFlagRequired("message-data")
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_generateMacCmd)
}
