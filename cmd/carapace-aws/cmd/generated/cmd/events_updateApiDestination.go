package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_updateApiDestinationCmd = &cobra.Command{
	Use:   "update-api-destination",
	Short: "Updates an API destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_updateApiDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_updateApiDestinationCmd).Standalone()

		events_updateApiDestinationCmd.Flags().String("connection-arn", "", "The ARN of the connection to use for the API destination.")
		events_updateApiDestinationCmd.Flags().String("description", "", "The name of the API destination to update.")
		events_updateApiDestinationCmd.Flags().String("http-method", "", "The method to use for the API destination.")
		events_updateApiDestinationCmd.Flags().String("invocation-endpoint", "", "The URL to the endpoint to use for the API destination.")
		events_updateApiDestinationCmd.Flags().String("invocation-rate-limit-per-second", "", "The maximum number of invocations per second to send to the API destination.")
		events_updateApiDestinationCmd.Flags().String("name", "", "The name of the API destination to update.")
		events_updateApiDestinationCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_updateApiDestinationCmd)
}
