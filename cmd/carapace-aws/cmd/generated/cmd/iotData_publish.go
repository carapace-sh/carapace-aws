package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotData_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes an MQTT message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotData_publishCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotData_publishCmd).Standalone()

		iotData_publishCmd.Flags().String("content-type", "", "A UTF-8 encoded string that describes the content of the publishing message.")
		iotData_publishCmd.Flags().String("correlation-data", "", "The base64-encoded binary data used by the sender of the request message to identify which request the response message is for when it's received.")
		iotData_publishCmd.Flags().String("message-expiry", "", "A user-defined integer value that represents the message expiry interval in seconds.")
		iotData_publishCmd.Flags().String("payload", "", "The message body.")
		iotData_publishCmd.Flags().String("payload-format-indicator", "", "An `Enum` string value that indicates whether the payload is formatted as UTF-8.")
		iotData_publishCmd.Flags().String("qos", "", "The Quality of Service (QoS) level.")
		iotData_publishCmd.Flags().String("response-topic", "", "A UTF-8 encoded string that's used as the topic name for a response message.")
		iotData_publishCmd.Flags().String("retain", "", "A Boolean value that determines whether to set the RETAIN flag when the message is published.")
		iotData_publishCmd.Flags().String("topic", "", "The name of the MQTT topic.")
		iotData_publishCmd.Flags().String("user-properties", "", "A JSON string that contains an array of JSON objects.")
		iotData_publishCmd.MarkFlagRequired("topic")
	})
	iotDataCmd.AddCommand(iotData_publishCmd)
}
