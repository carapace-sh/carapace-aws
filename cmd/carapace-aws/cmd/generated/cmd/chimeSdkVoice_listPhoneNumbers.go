package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listPhoneNumbersCmd = &cobra.Command{
	Use:   "list-phone-numbers",
	Short: "Lists the phone numbers for the specified Amazon Chime SDK account, Amazon Chime SDK user, Amazon Chime SDK Voice Connector, or Amazon Chime SDK Voice Connector group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listPhoneNumbersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_listPhoneNumbersCmd).Standalone()

		chimeSdkVoice_listPhoneNumbersCmd.Flags().String("filter-name", "", "The filter to limit the number of results.")
		chimeSdkVoice_listPhoneNumbersCmd.Flags().String("filter-value", "", "The filter value.")
		chimeSdkVoice_listPhoneNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chimeSdkVoice_listPhoneNumbersCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
		chimeSdkVoice_listPhoneNumbersCmd.Flags().String("product-type", "", "The phone number product types.")
		chimeSdkVoice_listPhoneNumbersCmd.Flags().String("status", "", "The status of your organization's phone numbers.")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listPhoneNumbersCmd)
}
