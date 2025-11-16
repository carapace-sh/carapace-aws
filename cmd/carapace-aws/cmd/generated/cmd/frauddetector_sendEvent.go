package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_sendEventCmd = &cobra.Command{
	Use:   "send-event",
	Short: "Stores events in Amazon Fraud Detector without generating fraud predictions for those events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_sendEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_sendEventCmd).Standalone()

		frauddetector_sendEventCmd.Flags().String("assigned-label", "", "The label to associate with the event.")
		frauddetector_sendEventCmd.Flags().String("entities", "", "An array of entities.")
		frauddetector_sendEventCmd.Flags().String("event-id", "", "The event ID to upload.")
		frauddetector_sendEventCmd.Flags().String("event-timestamp", "", "The timestamp that defines when the event under evaluation occurred.")
		frauddetector_sendEventCmd.Flags().String("event-type-name", "", "The event type name of the event.")
		frauddetector_sendEventCmd.Flags().String("event-variables", "", "Names of the event type's variables you defined in Amazon Fraud Detector to represent data elements and their corresponding values for the event you are sending for evaluation.")
		frauddetector_sendEventCmd.Flags().String("label-timestamp", "", "The timestamp associated with the label.")
		frauddetector_sendEventCmd.MarkFlagRequired("entities")
		frauddetector_sendEventCmd.MarkFlagRequired("event-id")
		frauddetector_sendEventCmd.MarkFlagRequired("event-timestamp")
		frauddetector_sendEventCmd.MarkFlagRequired("event-type-name")
		frauddetector_sendEventCmd.MarkFlagRequired("event-variables")
	})
	frauddetectorCmd.AddCommand(frauddetector_sendEventCmd)
}
