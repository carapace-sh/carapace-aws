package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeAccountAttributesCmd = &cobra.Command{
	Use:   "describe-account-attributes",
	Short: "Describes attributes of your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeAccountAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_describeAccountAttributesCmd).Standalone()

		pinpointSmsVoiceV2_describeAccountAttributesCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_describeAccountAttributesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeAccountAttributesCmd)
}
