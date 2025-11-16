package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listClassificationJobsCmd = &cobra.Command{
	Use:   "list-classification-jobs",
	Short: "Retrieves a subset of information about one or more classification jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listClassificationJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_listClassificationJobsCmd).Standalone()

		macie2_listClassificationJobsCmd.Flags().String("filter-criteria", "", "The criteria to use to filter the results.")
		macie2_listClassificationJobsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of the response.")
		macie2_listClassificationJobsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
		macie2_listClassificationJobsCmd.Flags().String("sort-criteria", "", "The criteria to use to sort the results.")
	})
	macie2Cmd.AddCommand(macie2_listClassificationJobsCmd)
}
