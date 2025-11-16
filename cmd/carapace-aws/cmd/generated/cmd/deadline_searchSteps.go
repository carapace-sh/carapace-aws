package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_searchStepsCmd = &cobra.Command{
	Use:   "search-steps",
	Short: "Searches for steps.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_searchStepsCmd).Standalone()

	deadline_searchStepsCmd.Flags().String("farm-id", "", "The farm ID to use for the step search.")
	deadline_searchStepsCmd.Flags().String("filter-expressions", "", "The filter expression, `AND` or `OR`, to use when searching among a group of search strings in a resource.")
	deadline_searchStepsCmd.Flags().String("item-offset", "", "Defines how far into the scrollable list to start the return of results.")
	deadline_searchStepsCmd.Flags().String("job-id", "", "The job ID to use in the step search.")
	deadline_searchStepsCmd.Flags().String("page-size", "", "Specifies the number of items per page for the resource.")
	deadline_searchStepsCmd.Flags().String("queue-ids", "", "The queue IDs in the step search.")
	deadline_searchStepsCmd.Flags().String("sort-expressions", "", "The search terms for a resource.")
	deadline_searchStepsCmd.MarkFlagRequired("farm-id")
	deadline_searchStepsCmd.MarkFlagRequired("item-offset")
	deadline_searchStepsCmd.MarkFlagRequired("queue-ids")
	deadlineCmd.AddCommand(deadline_searchStepsCmd)
}
