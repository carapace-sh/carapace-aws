package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeKeywordsCmd = &cobra.Command{
	Use:   "describe-keywords",
	Short: "Describes the specified keywords or all keywords on your origination phone number or pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeKeywordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_describeKeywordsCmd).Standalone()

		pinpointSmsVoiceV2_describeKeywordsCmd.Flags().String("filters", "", "An array of keyword filters to filter the results.")
		pinpointSmsVoiceV2_describeKeywordsCmd.Flags().String("keywords", "", "An array of keywords to search for.")
		pinpointSmsVoiceV2_describeKeywordsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_describeKeywordsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		pinpointSmsVoiceV2_describeKeywordsCmd.Flags().String("origination-identity", "", "The origination identity to use such as a PhoneNumberId, PhoneNumberArn, SenderId or SenderIdArn.")
		pinpointSmsVoiceV2_describeKeywordsCmd.MarkFlagRequired("origination-identity")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeKeywordsCmd)
}
