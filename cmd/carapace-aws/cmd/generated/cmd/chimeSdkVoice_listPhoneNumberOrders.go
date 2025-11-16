package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listPhoneNumberOrdersCmd = &cobra.Command{
	Use:   "list-phone-number-orders",
	Short: "Lists the phone numbers for an administrator's Amazon Chime SDK account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listPhoneNumberOrdersCmd).Standalone()

	chimeSdkVoice_listPhoneNumberOrdersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chimeSdkVoice_listPhoneNumberOrdersCmd.Flags().String("next-token", "", "The token used to retrieve the next page of results.")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listPhoneNumberOrdersCmd)
}
