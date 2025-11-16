package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_putEventTypeCmd = &cobra.Command{
	Use:   "put-event-type",
	Short: "Creates or updates an event type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_putEventTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_putEventTypeCmd).Standalone()

		frauddetector_putEventTypeCmd.Flags().String("description", "", "The description of the event type.")
		frauddetector_putEventTypeCmd.Flags().String("entity-types", "", "The entity type for the event type.")
		frauddetector_putEventTypeCmd.Flags().String("event-ingestion", "", "Specifies if ingestion is enabled or disabled.")
		frauddetector_putEventTypeCmd.Flags().String("event-orchestration", "", "Enables or disables event orchestration.")
		frauddetector_putEventTypeCmd.Flags().String("event-variables", "", "The event type variables.")
		frauddetector_putEventTypeCmd.Flags().String("labels", "", "The event type labels.")
		frauddetector_putEventTypeCmd.Flags().String("name", "", "The name.")
		frauddetector_putEventTypeCmd.Flags().String("tags", "", "A collection of key and value pairs.")
		frauddetector_putEventTypeCmd.MarkFlagRequired("entity-types")
		frauddetector_putEventTypeCmd.MarkFlagRequired("event-variables")
		frauddetector_putEventTypeCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_putEventTypeCmd)
}
