package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getFindingStatisticsCmd = &cobra.Command{
	Use:   "get-finding-statistics",
	Short: "Retrieves (queries) aggregated statistical data about findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getFindingStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getFindingStatisticsCmd).Standalone()

		macie2_getFindingStatisticsCmd.Flags().String("finding-criteria", "", "The criteria to use to filter the query results.")
		macie2_getFindingStatisticsCmd.Flags().String("group-by", "", "The finding property to use to group the query results.")
		macie2_getFindingStatisticsCmd.Flags().String("size", "", "The maximum number of items to include in each page of the response.")
		macie2_getFindingStatisticsCmd.Flags().String("sort-criteria", "", "The criteria to use to sort the query results.")
		macie2_getFindingStatisticsCmd.MarkFlagRequired("group-by")
	})
	macie2Cmd.AddCommand(macie2_getFindingStatisticsCmd)
}
