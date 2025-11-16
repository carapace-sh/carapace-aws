package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeReservedElasticsearchInstancesCmd = &cobra.Command{
	Use:   "describe-reserved-elasticsearch-instances",
	Short: "Returns information about reserved Elasticsearch instances for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeReservedElasticsearchInstancesCmd).Standalone()

	es_describeReservedElasticsearchInstancesCmd.Flags().String("max-results", "", "Set this value to limit the number of results returned.")
	es_describeReservedElasticsearchInstancesCmd.Flags().String("next-token", "", "NextToken should be sent in case if earlier API call produced result containing NextToken.")
	es_describeReservedElasticsearchInstancesCmd.Flags().String("reserved-elasticsearch-instance-id", "", "The reserved instance identifier filter value.")
	esCmd.AddCommand(es_describeReservedElasticsearchInstancesCmd)
}
