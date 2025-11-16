package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_createEventActionCmd = &cobra.Command{
	Use:   "create-event-action",
	Short: "This operation creates an event action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_createEventActionCmd).Standalone()

	dataexchange_createEventActionCmd.Flags().String("action", "", "What occurs after a certain event.")
	dataexchange_createEventActionCmd.Flags().String("event", "", "What occurs to start an action.")
	dataexchange_createEventActionCmd.Flags().String("tags", "", "Key-value pairs that you can associate with the event action.")
	dataexchange_createEventActionCmd.MarkFlagRequired("action")
	dataexchange_createEventActionCmd.MarkFlagRequired("event")
	dataexchangeCmd.AddCommand(dataexchange_createEventActionCmd)
}
