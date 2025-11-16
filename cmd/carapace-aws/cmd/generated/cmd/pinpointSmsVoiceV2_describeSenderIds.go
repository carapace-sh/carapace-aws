package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeSenderIdsCmd = &cobra.Command{
	Use:   "describe-sender-ids",
	Short: "Describes the specified SenderIds or all SenderIds associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeSenderIdsCmd).Standalone()

	pinpointSmsVoiceV2_describeSenderIdsCmd.Flags().String("filters", "", "An array of SenderIdFilter objects to filter the results.")
	pinpointSmsVoiceV2_describeSenderIdsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeSenderIdsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeSenderIdsCmd.Flags().String("owner", "", "Use `SELF` to filter the list of Sender Ids to ones your account owns or use `SHARED` to filter on Sender Ids shared with your account.")
	pinpointSmsVoiceV2_describeSenderIdsCmd.Flags().String("sender-ids", "", "An array of SenderIdAndCountry objects to search for.")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeSenderIdsCmd)
}
