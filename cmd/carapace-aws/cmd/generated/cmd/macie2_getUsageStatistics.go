package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getUsageStatisticsCmd = &cobra.Command{
	Use:   "get-usage-statistics",
	Short: "Retrieves (queries) quotas and aggregated usage data for one or more accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getUsageStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getUsageStatisticsCmd).Standalone()

		macie2_getUsageStatisticsCmd.Flags().String("filter-by", "", "An array of objects, one for each condition to use to filter the query results.")
		macie2_getUsageStatisticsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of the response.")
		macie2_getUsageStatisticsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
		macie2_getUsageStatisticsCmd.Flags().String("sort-by", "", "The criteria to use to sort the query results.")
		macie2_getUsageStatisticsCmd.Flags().String("time-range", "", "The inclusive time period to query usage data for.")
	})
	macie2Cmd.AddCommand(macie2_getUsageStatisticsCmd)
}
