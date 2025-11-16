package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describePhoneNumbersCmd = &cobra.Command{
	Use:   "describe-phone-numbers",
	Short: "Describes the specified origination phone number, or all the phone numbers in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describePhoneNumbersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_describePhoneNumbersCmd).Standalone()

		pinpointSmsVoiceV2_describePhoneNumbersCmd.Flags().String("filters", "", "An array of PhoneNumberFilter objects to filter the results.")
		pinpointSmsVoiceV2_describePhoneNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_describePhoneNumbersCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		pinpointSmsVoiceV2_describePhoneNumbersCmd.Flags().String("owner", "", "Use `SELF` to filter the list of phone numbers to ones your account owns or use `SHARED` to filter on phone numbers shared with your account.")
		pinpointSmsVoiceV2_describePhoneNumbersCmd.Flags().String("phone-number-ids", "", "The unique identifier of phone numbers to find information about.")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describePhoneNumbersCmd)
}
