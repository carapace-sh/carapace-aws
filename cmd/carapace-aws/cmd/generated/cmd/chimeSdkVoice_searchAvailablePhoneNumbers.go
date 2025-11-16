package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_searchAvailablePhoneNumbersCmd = &cobra.Command{
	Use:   "search-available-phone-numbers",
	Short: "Searches the provisioned phone numbers in an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_searchAvailablePhoneNumbersCmd).Standalone()

	chimeSdkVoice_searchAvailablePhoneNumbersCmd.Flags().String("area-code", "", "Confines a search to just the phone numbers associated with the specified area code.")
	chimeSdkVoice_searchAvailablePhoneNumbersCmd.Flags().String("city", "", "Confines a search to just the phone numbers associated with the specified city.")
	chimeSdkVoice_searchAvailablePhoneNumbersCmd.Flags().String("country", "", "Confines a search to just the phone numbers associated with the specified country.")
	chimeSdkVoice_searchAvailablePhoneNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	chimeSdkVoice_searchAvailablePhoneNumbersCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
	chimeSdkVoice_searchAvailablePhoneNumbersCmd.Flags().String("phone-number-type", "", "Confines a search to just the phone numbers associated with the specified phone number type, either **local** or **toll-free**.")
	chimeSdkVoice_searchAvailablePhoneNumbersCmd.Flags().String("state", "", "Confines a search to just the phone numbers associated with the specified state.")
	chimeSdkVoice_searchAvailablePhoneNumbersCmd.Flags().String("toll-free-prefix", "", "Confines a search to just the phone numbers associated with the specified toll-free prefix.")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_searchAvailablePhoneNumbersCmd)
}
