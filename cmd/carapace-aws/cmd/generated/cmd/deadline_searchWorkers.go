package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_searchWorkersCmd = &cobra.Command{
	Use:   "search-workers",
	Short: "Searches for workers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_searchWorkersCmd).Standalone()

	deadline_searchWorkersCmd.Flags().String("farm-id", "", "The farm ID in the workers search.")
	deadline_searchWorkersCmd.Flags().String("filter-expressions", "", "The filter expression, `AND` or `OR`, to use when searching among a group of search strings in a resource.")
	deadline_searchWorkersCmd.Flags().String("fleet-ids", "", "The fleet ID of the workers to search for.")
	deadline_searchWorkersCmd.Flags().String("item-offset", "", "Defines how far into the scrollable list to start the return of results.")
	deadline_searchWorkersCmd.Flags().String("page-size", "", "Specifies the number of items per page for the resource.")
	deadline_searchWorkersCmd.Flags().String("sort-expressions", "", "The search terms for a resource.")
	deadline_searchWorkersCmd.MarkFlagRequired("farm-id")
	deadline_searchWorkersCmd.MarkFlagRequired("fleet-ids")
	deadline_searchWorkersCmd.MarkFlagRequired("item-offset")
	deadlineCmd.AddCommand(deadline_searchWorkersCmd)
}
