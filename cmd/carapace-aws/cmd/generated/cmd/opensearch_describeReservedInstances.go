package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeReservedInstancesCmd = &cobra.Command{
	Use:   "describe-reserved-instances",
	Short: "Describes the Amazon OpenSearch Service instances that you have reserved in a given Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeReservedInstancesCmd).Standalone()

	opensearch_describeReservedInstancesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_describeReservedInstancesCmd.Flags().String("next-token", "", "If your initial `DescribeReservedInstances` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `DescribeReservedInstances` operations, which returns results in the next page.")
	opensearch_describeReservedInstancesCmd.Flags().String("reserved-instance-id", "", "The reserved instance identifier filter value.")
	opensearchCmd.AddCommand(opensearch_describeReservedInstancesCmd)
}
