package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateEventLabelCmd = &cobra.Command{
	Use:   "update-event-label",
	Short: "Updates the specified event with a new label.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateEventLabelCmd).Standalone()

	frauddetector_updateEventLabelCmd.Flags().String("assigned-label", "", "The new label to assign to the event.")
	frauddetector_updateEventLabelCmd.Flags().String("event-id", "", "The ID of the event associated with the label to update.")
	frauddetector_updateEventLabelCmd.Flags().String("event-type-name", "", "The event type of the event associated with the label to update.")
	frauddetector_updateEventLabelCmd.Flags().String("label-timestamp", "", "The timestamp associated with the label.")
	frauddetector_updateEventLabelCmd.MarkFlagRequired("assigned-label")
	frauddetector_updateEventLabelCmd.MarkFlagRequired("event-id")
	frauddetector_updateEventLabelCmd.MarkFlagRequired("event-type-name")
	frauddetector_updateEventLabelCmd.MarkFlagRequired("label-timestamp")
	frauddetectorCmd.AddCommand(frauddetector_updateEventLabelCmd)
}
