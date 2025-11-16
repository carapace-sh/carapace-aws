package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getFindingStatisticsV2Cmd = &cobra.Command{
	Use:   "get-finding-statistics-v2",
	Short: "Returns aggregated statistical data about findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getFindingStatisticsV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_getFindingStatisticsV2Cmd).Standalone()

		securityhub_getFindingStatisticsV2Cmd.Flags().String("group-by-rules", "", "Specifies how security findings should be aggregated and organized in the statistical analysis.")
		securityhub_getFindingStatisticsV2Cmd.Flags().String("max-statistic-results", "", "The maximum number of results to be returned.")
		securityhub_getFindingStatisticsV2Cmd.Flags().String("sort-order", "", "Orders the aggregation count in descending or ascending order.")
		securityhub_getFindingStatisticsV2Cmd.MarkFlagRequired("group-by-rules")
	})
	securityhubCmd.AddCommand(securityhub_getFindingStatisticsV2Cmd)
}
