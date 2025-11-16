package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeReservedNodesOfferingsCmd = &cobra.Command{
	Use:   "describe-reserved-nodes-offerings",
	Short: "Lists available reserved node offerings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeReservedNodesOfferingsCmd).Standalone()

	memorydb_describeReservedNodesOfferingsCmd.Flags().String("duration", "", "Duration filter value, specified in years or seconds.")
	memorydb_describeReservedNodesOfferingsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
	memorydb_describeReservedNodesOfferingsCmd.Flags().String("next-token", "", "An optional marker returned from a prior request.")
	memorydb_describeReservedNodesOfferingsCmd.Flags().String("node-type", "", "The node type for the reserved nodes.")
	memorydb_describeReservedNodesOfferingsCmd.Flags().String("offering-type", "", "The offering type filter value.")
	memorydb_describeReservedNodesOfferingsCmd.Flags().String("reserved-nodes-offering-id", "", "The offering identifier filter value.")
	memorydbCmd.AddCommand(memorydb_describeReservedNodesOfferingsCmd)
}
