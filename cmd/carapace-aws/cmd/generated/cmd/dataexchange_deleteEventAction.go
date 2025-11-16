package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_deleteEventActionCmd = &cobra.Command{
	Use:   "delete-event-action",
	Short: "This operation deletes the event action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_deleteEventActionCmd).Standalone()

	dataexchange_deleteEventActionCmd.Flags().String("event-action-id", "", "The unique identifier for the event action.")
	dataexchange_deleteEventActionCmd.MarkFlagRequired("event-action-id")
	dataexchangeCmd.AddCommand(dataexchange_deleteEventActionCmd)
}
