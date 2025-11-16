package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd = &cobra.Command{
	Use:   "describe-verified-destination-numbers",
	Short: "Retrieves the specified verified destination numbers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd).Standalone()

		pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd.Flags().String("destination-phone-numbers", "", "An array of verified destination phone number, in E.164 format.")
		pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd.Flags().String("filters", "", "An array of VerifiedDestinationNumberFilter objects to filter the results.")
		pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd.Flags().String("verified-destination-number-ids", "", "An array of VerifiedDestinationNumberid to retrieve.")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeVerifiedDestinationNumbersCmd)
}
