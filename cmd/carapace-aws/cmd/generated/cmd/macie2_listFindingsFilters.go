package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listFindingsFiltersCmd = &cobra.Command{
	Use:   "list-findings-filters",
	Short: "Retrieves a subset of information about all the findings filters for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listFindingsFiltersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_listFindingsFiltersCmd).Standalone()

		macie2_listFindingsFiltersCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of a paginated response.")
		macie2_listFindingsFiltersCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	})
	macie2Cmd.AddCommand(macie2_listFindingsFiltersCmd)
}
