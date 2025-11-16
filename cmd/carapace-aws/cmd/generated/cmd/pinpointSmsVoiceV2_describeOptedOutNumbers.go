package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeOptedOutNumbersCmd = &cobra.Command{
	Use:   "describe-opted-out-numbers",
	Short: "Describes the specified opted out destination numbers or all opted out destination numbers in an opt-out list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeOptedOutNumbersCmd).Standalone()

	pinpointSmsVoiceV2_describeOptedOutNumbersCmd.Flags().String("filters", "", "An array of OptedOutFilter objects to filter the results on.")
	pinpointSmsVoiceV2_describeOptedOutNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeOptedOutNumbersCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeOptedOutNumbersCmd.Flags().String("opt-out-list-name", "", "The OptOutListName or OptOutListArn of the OptOutList.")
	pinpointSmsVoiceV2_describeOptedOutNumbersCmd.Flags().String("opted-out-numbers", "", "An array of phone numbers to search for in the OptOutList.")
	pinpointSmsVoiceV2_describeOptedOutNumbersCmd.MarkFlagRequired("opt-out-list-name")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeOptedOutNumbersCmd)
}
