package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getBucketsAggregationCmd = &cobra.Command{
	Use:   "get-buckets-aggregation",
	Short: "Aggregates on indexed data with search queries pertaining to particular fields.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getBucketsAggregationCmd).Standalone()

	iot_getBucketsAggregationCmd.Flags().String("aggregation-field", "", "The aggregation field.")
	iot_getBucketsAggregationCmd.Flags().String("buckets-aggregation-type", "", "The basic control of the response shape and the bucket aggregation type to perform.")
	iot_getBucketsAggregationCmd.Flags().String("index-name", "", "The name of the index to search.")
	iot_getBucketsAggregationCmd.Flags().String("query-string", "", "The search query string.")
	iot_getBucketsAggregationCmd.Flags().String("query-version", "", "The version of the query.")
	iot_getBucketsAggregationCmd.MarkFlagRequired("aggregation-field")
	iot_getBucketsAggregationCmd.MarkFlagRequired("buckets-aggregation-type")
	iot_getBucketsAggregationCmd.MarkFlagRequired("query-string")
	iotCmd.AddCommand(iot_getBucketsAggregationCmd)
}
