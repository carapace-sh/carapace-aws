package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getEventPredictionCmd = &cobra.Command{
	Use:   "get-event-prediction",
	Short: "Evaluates an event against a detector version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getEventPredictionCmd).Standalone()

	frauddetector_getEventPredictionCmd.Flags().String("detector-id", "", "The detector ID.")
	frauddetector_getEventPredictionCmd.Flags().String("detector-version-id", "", "The detector version ID.")
	frauddetector_getEventPredictionCmd.Flags().String("entities", "", "The entity type (associated with the detector's event type) and specific entity ID representing who performed the event.")
	frauddetector_getEventPredictionCmd.Flags().String("event-id", "", "The unique ID used to identify the event.")
	frauddetector_getEventPredictionCmd.Flags().String("event-timestamp", "", "Timestamp that defines when the event under evaluation occurred.")
	frauddetector_getEventPredictionCmd.Flags().String("event-type-name", "", "The event type associated with the detector specified for the prediction.")
	frauddetector_getEventPredictionCmd.Flags().String("event-variables", "", "Names of the event type's variables you defined in Amazon Fraud Detector to represent data elements and their corresponding values for the event you are sending for evaluation.")
	frauddetector_getEventPredictionCmd.Flags().String("external-model-endpoint-data-blobs", "", "The Amazon SageMaker model endpoint input data blobs.")
	frauddetector_getEventPredictionCmd.MarkFlagRequired("detector-id")
	frauddetector_getEventPredictionCmd.MarkFlagRequired("entities")
	frauddetector_getEventPredictionCmd.MarkFlagRequired("event-id")
	frauddetector_getEventPredictionCmd.MarkFlagRequired("event-timestamp")
	frauddetector_getEventPredictionCmd.MarkFlagRequired("event-type-name")
	frauddetector_getEventPredictionCmd.MarkFlagRequired("event-variables")
	frauddetectorCmd.AddCommand(frauddetector_getEventPredictionCmd)
}
