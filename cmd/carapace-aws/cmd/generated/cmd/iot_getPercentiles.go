package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getPercentilesCmd = &cobra.Command{
	Use:   "get-percentiles",
	Short: "Groups the aggregated values that match the query into percentile groupings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getPercentilesCmd).Standalone()

	iot_getPercentilesCmd.Flags().String("aggregation-field", "", "The field to aggregate.")
	iot_getPercentilesCmd.Flags().String("index-name", "", "The name of the index to search.")
	iot_getPercentilesCmd.Flags().String("percents", "", "The percentile groups returned.")
	iot_getPercentilesCmd.Flags().String("query-string", "", "The search query string.")
	iot_getPercentilesCmd.Flags().String("query-version", "", "The query version.")
	iot_getPercentilesCmd.MarkFlagRequired("query-string")
	iotCmd.AddCommand(iot_getPercentilesCmd)
}
