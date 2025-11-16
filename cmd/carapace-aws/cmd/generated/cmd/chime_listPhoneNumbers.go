package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_listPhoneNumbersCmd = &cobra.Command{
	Use:   "list-phone-numbers",
	Short: "Lists the phone numbers for the specified Amazon Chime account, Amazon Chime user, Amazon Chime Voice Connector, or Amazon Chime Voice Connector group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_listPhoneNumbersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_listPhoneNumbersCmd).Standalone()

		chime_listPhoneNumbersCmd.Flags().String("filter-name", "", "The filter to use to limit the number of results.")
		chime_listPhoneNumbersCmd.Flags().String("filter-value", "", "The value to use for the filter.")
		chime_listPhoneNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chime_listPhoneNumbersCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		chime_listPhoneNumbersCmd.Flags().String("product-type", "", "The phone number product type.")
		chime_listPhoneNumbersCmd.Flags().String("status", "", "The phone number status.")
	})
	chimeCmd.AddCommand(chime_listPhoneNumbersCmd)
}
