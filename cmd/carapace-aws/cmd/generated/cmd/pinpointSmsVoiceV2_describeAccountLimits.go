package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeAccountLimitsCmd = &cobra.Command{
	Use:   "describe-account-limits",
	Short: "Describes the current End User MessagingSMS SMS Voice V2 resource quotas for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeAccountLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_describeAccountLimitsCmd).Standalone()

		pinpointSmsVoiceV2_describeAccountLimitsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_describeAccountLimitsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeAccountLimitsCmd)
}
