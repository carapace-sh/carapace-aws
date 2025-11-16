package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_createApiDestinationCmd = &cobra.Command{
	Use:   "create-api-destination",
	Short: "Creates an API destination, which is an HTTP invocation endpoint configured as a target for events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_createApiDestinationCmd).Standalone()

	events_createApiDestinationCmd.Flags().String("connection-arn", "", "The ARN of the connection to use for the API destination.")
	events_createApiDestinationCmd.Flags().String("description", "", "A description for the API destination to create.")
	events_createApiDestinationCmd.Flags().String("http-method", "", "The method to use for the request to the HTTP invocation endpoint.")
	events_createApiDestinationCmd.Flags().String("invocation-endpoint", "", "The URL to the HTTP invocation endpoint for the API destination.")
	events_createApiDestinationCmd.Flags().String("invocation-rate-limit-per-second", "", "The maximum number of requests per second to send to the HTTP invocation endpoint.")
	events_createApiDestinationCmd.Flags().String("name", "", "The name for the API destination to create.")
	events_createApiDestinationCmd.MarkFlagRequired("connection-arn")
	events_createApiDestinationCmd.MarkFlagRequired("http-method")
	events_createApiDestinationCmd.MarkFlagRequired("invocation-endpoint")
	events_createApiDestinationCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_createApiDestinationCmd)
}
