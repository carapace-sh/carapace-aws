package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listFiltersCmd = &cobra.Command{
	Use:   "list-filters",
	Short: "Lists all filters that belong to a given dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listFiltersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_listFiltersCmd).Standalone()

		personalize_listFiltersCmd.Flags().String("dataset-group-arn", "", "The ARN of the dataset group that contains the filters.")
		personalize_listFiltersCmd.Flags().String("max-results", "", "The maximum number of filters to return.")
		personalize_listFiltersCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListFilters` for getting the next set of filters (if they exist).")
	})
	personalizeCmd.AddCommand(personalize_listFiltersCmd)
}
