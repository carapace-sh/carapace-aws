package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_generateMacEmvPinChangeCmd = &cobra.Command{
	Use:   "generate-mac-emv-pin-change",
	Short: "Generates an issuer script mac for EMV payment cards that use offline PINs as the cardholder verification method (CVM).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_generateMacEmvPinChangeCmd).Standalone()

	paymentCryptographyData_generateMacEmvPinChangeCmd.Flags().String("derivation-method-attributes", "", "The attributes and data values to derive payment card specific confidentiality and integrity keys.")
	paymentCryptographyData_generateMacEmvPinChangeCmd.Flags().String("message-data", "", "The message data is the APDU command from the card reader or terminal.")
	paymentCryptographyData_generateMacEmvPinChangeCmd.Flags().String("new-encrypted-pin-block", "", "The incoming new encrypted PIN block data for offline pin change on an EMV card.")
	paymentCryptographyData_generateMacEmvPinChangeCmd.Flags().String("new-pin-pek-identifier", "", "The `keyARN` of the PEK protecting the incoming new encrypted PIN block.")
	paymentCryptographyData_generateMacEmvPinChangeCmd.Flags().String("pin-block-format", "", "The PIN encoding format of the incoming new encrypted PIN block as specified in ISO 9564.")
	paymentCryptographyData_generateMacEmvPinChangeCmd.Flags().String("secure-messaging-confidentiality-key-identifier", "", "The `keyARN` of the issuer master key (IMK-SMC) used to protect the PIN block data in the issuer script response.")
	paymentCryptographyData_generateMacEmvPinChangeCmd.Flags().String("secure-messaging-integrity-key-identifier", "", "The `keyARN` of the issuer master key (IMK-SMI) used to authenticate the issuer script response.")
	paymentCryptographyData_generateMacEmvPinChangeCmd.MarkFlagRequired("derivation-method-attributes")
	paymentCryptographyData_generateMacEmvPinChangeCmd.MarkFlagRequired("message-data")
	paymentCryptographyData_generateMacEmvPinChangeCmd.MarkFlagRequired("new-encrypted-pin-block")
	paymentCryptographyData_generateMacEmvPinChangeCmd.MarkFlagRequired("new-pin-pek-identifier")
	paymentCryptographyData_generateMacEmvPinChangeCmd.MarkFlagRequired("pin-block-format")
	paymentCryptographyData_generateMacEmvPinChangeCmd.MarkFlagRequired("secure-messaging-confidentiality-key-identifier")
	paymentCryptographyData_generateMacEmvPinChangeCmd.MarkFlagRequired("secure-messaging-integrity-key-identifier")
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_generateMacEmvPinChangeCmd)
}
