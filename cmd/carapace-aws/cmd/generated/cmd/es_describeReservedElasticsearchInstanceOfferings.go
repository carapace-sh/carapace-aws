package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeReservedElasticsearchInstanceOfferingsCmd = &cobra.Command{
	Use:   "describe-reserved-elasticsearch-instance-offerings",
	Short: "Lists available reserved Elasticsearch instance offerings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeReservedElasticsearchInstanceOfferingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_describeReservedElasticsearchInstanceOfferingsCmd).Standalone()

		es_describeReservedElasticsearchInstanceOfferingsCmd.Flags().String("max-results", "", "Set this value to limit the number of results returned.")
		es_describeReservedElasticsearchInstanceOfferingsCmd.Flags().String("next-token", "", "NextToken should be sent in case if earlier API call produced result containing NextToken.")
		es_describeReservedElasticsearchInstanceOfferingsCmd.Flags().String("reserved-elasticsearch-instance-offering-id", "", "The offering identifier filter value.")
	})
	esCmd.AddCommand(es_describeReservedElasticsearchInstanceOfferingsCmd)
}
