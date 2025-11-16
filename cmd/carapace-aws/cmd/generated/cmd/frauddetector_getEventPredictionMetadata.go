package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getEventPredictionMetadataCmd = &cobra.Command{
	Use:   "get-event-prediction-metadata",
	Short: "Gets details of the past fraud predictions for the specified event ID, event type, detector ID, and detector version ID that was generated in the specified time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getEventPredictionMetadataCmd).Standalone()

	frauddetector_getEventPredictionMetadataCmd.Flags().String("detector-id", "", "The detector ID.")
	frauddetector_getEventPredictionMetadataCmd.Flags().String("detector-version-id", "", "The detector version ID.")
	frauddetector_getEventPredictionMetadataCmd.Flags().String("event-id", "", "The event ID.")
	frauddetector_getEventPredictionMetadataCmd.Flags().String("event-type-name", "", "The event type associated with the detector specified for the prediction.")
	frauddetector_getEventPredictionMetadataCmd.Flags().String("prediction-timestamp", "", "The timestamp that defines when the prediction was generated.")
	frauddetector_getEventPredictionMetadataCmd.MarkFlagRequired("detector-id")
	frauddetector_getEventPredictionMetadataCmd.MarkFlagRequired("detector-version-id")
	frauddetector_getEventPredictionMetadataCmd.MarkFlagRequired("event-id")
	frauddetector_getEventPredictionMetadataCmd.MarkFlagRequired("event-type-name")
	frauddetector_getEventPredictionMetadataCmd.MarkFlagRequired("prediction-timestamp")
	frauddetectorCmd.AddCommand(frauddetector_getEventPredictionMetadataCmd)
}
