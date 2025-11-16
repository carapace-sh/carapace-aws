package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_updateEventDestinationCmd = &cobra.Command{
	Use:   "update-event-destination",
	Short: "Updates an existing event destination in a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_updateEventDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_updateEventDestinationCmd).Standalone()

		pinpointSmsVoiceV2_updateEventDestinationCmd.Flags().String("cloud-watch-logs-destination", "", "An object that contains information about an event destination that sends data to CloudWatch Logs.")
		pinpointSmsVoiceV2_updateEventDestinationCmd.Flags().String("configuration-set-name", "", "The configuration set to update with the new event destination.")
		pinpointSmsVoiceV2_updateEventDestinationCmd.Flags().Bool("enabled", false, "When set to true logging is enabled.")
		pinpointSmsVoiceV2_updateEventDestinationCmd.Flags().String("event-destination-name", "", "The name to use for the event destination.")
		pinpointSmsVoiceV2_updateEventDestinationCmd.Flags().String("kinesis-firehose-destination", "", "An object that contains information about an event destination for logging to Firehose.")
		pinpointSmsVoiceV2_updateEventDestinationCmd.Flags().String("matching-event-types", "", "An array of event types that determine which events to log.")
		pinpointSmsVoiceV2_updateEventDestinationCmd.Flags().Bool("no-enabled", false, "When set to true logging is enabled.")
		pinpointSmsVoiceV2_updateEventDestinationCmd.Flags().String("sns-destination", "", "An object that contains information about an event destination that sends data to Amazon SNS.")
		pinpointSmsVoiceV2_updateEventDestinationCmd.MarkFlagRequired("configuration-set-name")
		pinpointSmsVoiceV2_updateEventDestinationCmd.MarkFlagRequired("event-destination-name")
		pinpointSmsVoiceV2_updateEventDestinationCmd.Flag("no-enabled").Hidden = true
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_updateEventDestinationCmd)
}
