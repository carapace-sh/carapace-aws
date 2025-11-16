package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeSpendLimitsCmd = &cobra.Command{
	Use:   "describe-spend-limits",
	Short: "Describes the current monthly spend limits for sending voice and text messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeSpendLimitsCmd).Standalone()

	pinpointSmsVoiceV2_describeSpendLimitsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeSpendLimitsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeSpendLimitsCmd)
}
