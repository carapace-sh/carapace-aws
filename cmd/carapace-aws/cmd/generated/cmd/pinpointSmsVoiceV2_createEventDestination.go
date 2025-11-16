package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createEventDestinationCmd = &cobra.Command{
	Use:   "create-event-destination",
	Short: "Creates a new event destination in a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createEventDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_createEventDestinationCmd).Standalone()

		pinpointSmsVoiceV2_createEventDestinationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pinpointSmsVoiceV2_createEventDestinationCmd.Flags().String("cloud-watch-logs-destination", "", "An object that contains information about an event destination for logging to Amazon CloudWatch Logs.")
		pinpointSmsVoiceV2_createEventDestinationCmd.Flags().String("configuration-set-name", "", "Either the name of the configuration set or the configuration set ARN to apply event logging to.")
		pinpointSmsVoiceV2_createEventDestinationCmd.Flags().String("event-destination-name", "", "The name that identifies the event destination.")
		pinpointSmsVoiceV2_createEventDestinationCmd.Flags().String("kinesis-firehose-destination", "", "An object that contains information about an event destination for logging to Amazon Data Firehose.")
		pinpointSmsVoiceV2_createEventDestinationCmd.Flags().String("matching-event-types", "", "An array of event types that determine which events to log.")
		pinpointSmsVoiceV2_createEventDestinationCmd.Flags().String("sns-destination", "", "An object that contains information about an event destination for logging to Amazon SNS.")
		pinpointSmsVoiceV2_createEventDestinationCmd.MarkFlagRequired("configuration-set-name")
		pinpointSmsVoiceV2_createEventDestinationCmd.MarkFlagRequired("event-destination-name")
		pinpointSmsVoiceV2_createEventDestinationCmd.MarkFlagRequired("matching-event-types")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createEventDestinationCmd)
}
