package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_translatePinDataCmd = &cobra.Command{
	Use:   "translate-pin-data",
	Short: "Translates encrypted PIN block from and to ISO 9564 formats 0,1,3,4.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_translatePinDataCmd).Standalone()

	paymentCryptographyData_translatePinDataCmd.Flags().String("encrypted-pin-block", "", "The encrypted PIN block data that Amazon Web Services Payment Cryptography translates.")
	paymentCryptographyData_translatePinDataCmd.Flags().String("incoming-dukpt-attributes", "", "The attributes and values to use for incoming DUKPT encryption key for PIN block translation.")
	paymentCryptographyData_translatePinDataCmd.Flags().String("incoming-key-identifier", "", "The `keyARN` of the encryption key under which incoming PIN block data is encrypted.")
	paymentCryptographyData_translatePinDataCmd.Flags().String("incoming-translation-attributes", "", "The format of the incoming PIN block data for translation within Amazon Web Services Payment Cryptography.")
	paymentCryptographyData_translatePinDataCmd.Flags().String("incoming-wrapped-key", "", "The WrappedKeyBlock containing the encryption key under which incoming PIN block data is encrypted.")
	paymentCryptographyData_translatePinDataCmd.Flags().String("outgoing-dukpt-attributes", "", "The attributes and values to use for outgoing DUKPT encryption key after PIN block translation.")
	paymentCryptographyData_translatePinDataCmd.Flags().String("outgoing-key-identifier", "", "The `keyARN` of the encryption key for encrypting outgoing PIN block data.")
	paymentCryptographyData_translatePinDataCmd.Flags().String("outgoing-translation-attributes", "", "The format of the outgoing PIN block data after translation by Amazon Web Services Payment Cryptography.")
	paymentCryptographyData_translatePinDataCmd.Flags().String("outgoing-wrapped-key", "", "The WrappedKeyBlock containing the encryption key for encrypting outgoing PIN block data.")
	paymentCryptographyData_translatePinDataCmd.MarkFlagRequired("encrypted-pin-block")
	paymentCryptographyData_translatePinDataCmd.MarkFlagRequired("incoming-key-identifier")
	paymentCryptographyData_translatePinDataCmd.MarkFlagRequired("incoming-translation-attributes")
	paymentCryptographyData_translatePinDataCmd.MarkFlagRequired("outgoing-key-identifier")
	paymentCryptographyData_translatePinDataCmd.MarkFlagRequired("outgoing-translation-attributes")
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_translatePinDataCmd)
}
