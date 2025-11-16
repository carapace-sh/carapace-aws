package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_listScheduledQueriesCmd = &cobra.Command{
	Use:   "list-scheduled-queries",
	Short: "Gets a list of all scheduled queries in the caller's Amazon account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_listScheduledQueriesCmd).Standalone()

	timestreamQuery_listScheduledQueriesCmd.Flags().String("max-results", "", "The maximum number of items to return in the output.")
	timestreamQuery_listScheduledQueriesCmd.Flags().String("next-token", "", "A pagination token to resume pagination.")
	timestreamQueryCmd.AddCommand(timestreamQuery_listScheduledQueriesCmd)
}
