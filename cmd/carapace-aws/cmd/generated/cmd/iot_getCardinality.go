package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getCardinalityCmd = &cobra.Command{
	Use:   "get-cardinality",
	Short: "Returns the approximate count of unique values that match the query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getCardinalityCmd).Standalone()

	iot_getCardinalityCmd.Flags().String("aggregation-field", "", "The field to aggregate.")
	iot_getCardinalityCmd.Flags().String("index-name", "", "The name of the index to search.")
	iot_getCardinalityCmd.Flags().String("query-string", "", "The search query string.")
	iot_getCardinalityCmd.Flags().String("query-version", "", "The query version.")
	iot_getCardinalityCmd.MarkFlagRequired("query-string")
	iotCmd.AddCommand(iot_getCardinalityCmd)
}
