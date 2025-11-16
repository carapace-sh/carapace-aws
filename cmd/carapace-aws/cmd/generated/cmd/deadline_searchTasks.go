package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_searchTasksCmd = &cobra.Command{
	Use:   "search-tasks",
	Short: "Searches for tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_searchTasksCmd).Standalone()

	deadline_searchTasksCmd.Flags().String("farm-id", "", "The farm ID of the task.")
	deadline_searchTasksCmd.Flags().String("filter-expressions", "", "The filter expression, `AND` or `OR`, to use when searching among a group of search strings in a resource.")
	deadline_searchTasksCmd.Flags().String("item-offset", "", "Defines how far into the scrollable list to start the return of results.")
	deadline_searchTasksCmd.Flags().String("job-id", "", "The job ID for the task search.")
	deadline_searchTasksCmd.Flags().String("page-size", "", "Specifies the number of items per page for the resource.")
	deadline_searchTasksCmd.Flags().String("queue-ids", "", "The queue IDs to include in the search.")
	deadline_searchTasksCmd.Flags().String("sort-expressions", "", "The search terms for a resource.")
	deadline_searchTasksCmd.MarkFlagRequired("farm-id")
	deadline_searchTasksCmd.MarkFlagRequired("item-offset")
	deadline_searchTasksCmd.MarkFlagRequired("queue-ids")
	deadlineCmd.AddCommand(deadline_searchTasksCmd)
}
