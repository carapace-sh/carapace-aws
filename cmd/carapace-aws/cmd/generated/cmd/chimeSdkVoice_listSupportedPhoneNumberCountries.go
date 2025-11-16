package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listSupportedPhoneNumberCountriesCmd = &cobra.Command{
	Use:   "list-supported-phone-number-countries",
	Short: "Lists the countries that you can order phone numbers from.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listSupportedPhoneNumberCountriesCmd).Standalone()

	chimeSdkVoice_listSupportedPhoneNumberCountriesCmd.Flags().String("product-type", "", "The phone number product type.")
	chimeSdkVoice_listSupportedPhoneNumberCountriesCmd.MarkFlagRequired("product-type")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listSupportedPhoneNumberCountriesCmd)
}
