package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getStatisticsCmd = &cobra.Command{
	Use:   "get-statistics",
	Short: "Returns the count, average, sum, minimum, maximum, sum of squares, variance, and standard deviation for the specified aggregated field.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_getStatisticsCmd).Standalone()

		iot_getStatisticsCmd.Flags().String("aggregation-field", "", "The aggregation field name.")
		iot_getStatisticsCmd.Flags().String("index-name", "", "The name of the index to search.")
		iot_getStatisticsCmd.Flags().String("query-string", "", "The query used to search.")
		iot_getStatisticsCmd.Flags().String("query-version", "", "The version of the query used to search.")
		iot_getStatisticsCmd.MarkFlagRequired("query-string")
	})
	iotCmd.AddCommand(iot_getStatisticsCmd)
}
