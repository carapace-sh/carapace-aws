package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_searchJobsCmd = &cobra.Command{
	Use:   "search-jobs",
	Short: "Searches for jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_searchJobsCmd).Standalone()

	deadline_searchJobsCmd.Flags().String("farm-id", "", "The farm ID of the job.")
	deadline_searchJobsCmd.Flags().String("filter-expressions", "", "The filter expression, `AND` or `OR`, to use when searching among a group of search strings in a resource.")
	deadline_searchJobsCmd.Flags().String("item-offset", "", "Defines how far into the scrollable list to start the return of results.")
	deadline_searchJobsCmd.Flags().String("page-size", "", "Specifies the number of items per page for the resource.")
	deadline_searchJobsCmd.Flags().String("queue-ids", "", "The queue ID to use in the job search.")
	deadline_searchJobsCmd.Flags().String("sort-expressions", "", "The search terms for a resource.")
	deadline_searchJobsCmd.MarkFlagRequired("farm-id")
	deadline_searchJobsCmd.MarkFlagRequired("item-offset")
	deadline_searchJobsCmd.MarkFlagRequired("queue-ids")
	deadlineCmd.AddCommand(deadline_searchJobsCmd)
}
