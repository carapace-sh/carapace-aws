package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeReservedInstanceOfferingsCmd = &cobra.Command{
	Use:   "describe-reserved-instance-offerings",
	Short: "Describes the available Amazon OpenSearch Service Reserved Instance offerings for a given Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeReservedInstanceOfferingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_describeReservedInstanceOfferingsCmd).Standalone()

		opensearch_describeReservedInstanceOfferingsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearch_describeReservedInstanceOfferingsCmd.Flags().String("next-token", "", "If your initial `DescribeReservedInstanceOfferings` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `DescribeReservedInstanceOfferings` operations, which returns results in the next page.")
		opensearch_describeReservedInstanceOfferingsCmd.Flags().String("reserved-instance-offering-id", "", "The Reserved Instance identifier filter value.")
	})
	opensearchCmd.AddCommand(opensearch_describeReservedInstanceOfferingsCmd)
}
