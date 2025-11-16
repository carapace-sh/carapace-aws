package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_getEventActionCmd = &cobra.Command{
	Use:   "get-event-action",
	Short: "This operation retrieves information about an event action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_getEventActionCmd).Standalone()

	dataexchange_getEventActionCmd.Flags().String("event-action-id", "", "The unique identifier for the event action.")
	dataexchange_getEventActionCmd.MarkFlagRequired("event-action-id")
	dataexchangeCmd.AddCommand(dataexchange_getEventActionCmd)
}
