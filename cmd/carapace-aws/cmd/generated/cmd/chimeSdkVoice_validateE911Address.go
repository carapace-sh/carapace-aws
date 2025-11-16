package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_validateE911AddressCmd = &cobra.Command{
	Use:   "validate-e911-address",
	Short: "Validates an address to be used for 911 calls made with Amazon Chime SDK Voice Connectors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_validateE911AddressCmd).Standalone()

	chimeSdkVoice_validateE911AddressCmd.Flags().String("aws-account-id", "", "The AWS account ID.")
	chimeSdkVoice_validateE911AddressCmd.Flags().String("city", "", "The address city, such as `Portland`.")
	chimeSdkVoice_validateE911AddressCmd.Flags().String("country", "", "The country in the address being validated as two-letter country code in ISO 3166-1 alpha-2 format, such as `US`.")
	chimeSdkVoice_validateE911AddressCmd.Flags().String("postal-code", "", "The dress postal code, such `04352`.")
	chimeSdkVoice_validateE911AddressCmd.Flags().String("state", "", "The address state, such as `ME`.")
	chimeSdkVoice_validateE911AddressCmd.Flags().String("street-info", "", "The address street information, such as `8th Avenue`.")
	chimeSdkVoice_validateE911AddressCmd.Flags().String("street-number", "", "The address street number, such as `200` or `2121`.")
	chimeSdkVoice_validateE911AddressCmd.MarkFlagRequired("aws-account-id")
	chimeSdkVoice_validateE911AddressCmd.MarkFlagRequired("city")
	chimeSdkVoice_validateE911AddressCmd.MarkFlagRequired("country")
	chimeSdkVoice_validateE911AddressCmd.MarkFlagRequired("postal-code")
	chimeSdkVoice_validateE911AddressCmd.MarkFlagRequired("state")
	chimeSdkVoice_validateE911AddressCmd.MarkFlagRequired("street-info")
	chimeSdkVoice_validateE911AddressCmd.MarkFlagRequired("street-number")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_validateE911AddressCmd)
}
