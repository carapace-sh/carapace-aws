package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_searchAvailablePhoneNumbersCmd = &cobra.Command{
	Use:   "search-available-phone-numbers",
	Short: "Searches for phone numbers that can be ordered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_searchAvailablePhoneNumbersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_searchAvailablePhoneNumbersCmd).Standalone()

		chime_searchAvailablePhoneNumbersCmd.Flags().String("area-code", "", "The area code used to filter results.")
		chime_searchAvailablePhoneNumbersCmd.Flags().String("city", "", "The city used to filter results.")
		chime_searchAvailablePhoneNumbersCmd.Flags().String("country", "", "The country used to filter results.")
		chime_searchAvailablePhoneNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chime_searchAvailablePhoneNumbersCmd.Flags().String("next-token", "", "The token used to retrieve the next page of results.")
		chime_searchAvailablePhoneNumbersCmd.Flags().String("phone-number-type", "", "The phone number type used to filter results.")
		chime_searchAvailablePhoneNumbersCmd.Flags().String("state", "", "The state used to filter results.")
		chime_searchAvailablePhoneNumbersCmd.Flags().String("toll-free-prefix", "", "The toll-free prefix that you use to filter results.")
	})
	chimeCmd.AddCommand(chime_searchAvailablePhoneNumbersCmd)
}
