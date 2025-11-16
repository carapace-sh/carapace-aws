package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_updateEventActionCmd = &cobra.Command{
	Use:   "update-event-action",
	Short: "This operation updates the event action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_updateEventActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_updateEventActionCmd).Standalone()

		dataexchange_updateEventActionCmd.Flags().String("action", "", "What occurs after a certain event.")
		dataexchange_updateEventActionCmd.Flags().String("event-action-id", "", "The unique identifier for the event action.")
		dataexchange_updateEventActionCmd.MarkFlagRequired("event-action-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_updateEventActionCmd)
}
