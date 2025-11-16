package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_listEventActionsCmd = &cobra.Command{
	Use:   "list-event-actions",
	Short: "This operation lists your event actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_listEventActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_listEventActionsCmd).Standalone()

		dataexchange_listEventActionsCmd.Flags().String("event-source-id", "", "The unique identifier for the event source.")
		dataexchange_listEventActionsCmd.Flags().String("max-results", "", "The maximum number of results returned by a single call.")
		dataexchange_listEventActionsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	})
	dataexchangeCmd.AddCommand(dataexchange_listEventActionsCmd)
}
