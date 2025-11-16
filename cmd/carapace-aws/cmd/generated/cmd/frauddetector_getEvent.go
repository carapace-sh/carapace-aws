package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getEventCmd = &cobra.Command{
	Use:   "get-event",
	Short: "Retrieves details of events stored with Amazon Fraud Detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getEventCmd).Standalone()

		frauddetector_getEventCmd.Flags().String("event-id", "", "The ID of the event to retrieve.")
		frauddetector_getEventCmd.Flags().String("event-type-name", "", "The event type of the event to retrieve.")
		frauddetector_getEventCmd.MarkFlagRequired("event-id")
		frauddetector_getEventCmd.MarkFlagRequired("event-type-name")
	})
	frauddetectorCmd.AddCommand(frauddetector_getEventCmd)
}
