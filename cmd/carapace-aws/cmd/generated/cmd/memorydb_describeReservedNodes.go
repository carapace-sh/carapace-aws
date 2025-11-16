package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeReservedNodesCmd = &cobra.Command{
	Use:   "describe-reserved-nodes",
	Short: "Returns information about reserved nodes for this account, or about a specified reserved node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeReservedNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_describeReservedNodesCmd).Standalone()

		memorydb_describeReservedNodesCmd.Flags().String("duration", "", "The duration filter value, specified in years or seconds.")
		memorydb_describeReservedNodesCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		memorydb_describeReservedNodesCmd.Flags().String("next-token", "", "An optional marker returned from a prior request.")
		memorydb_describeReservedNodesCmd.Flags().String("node-type", "", "The node type filter value.")
		memorydb_describeReservedNodesCmd.Flags().String("offering-type", "", "The offering type filter value.")
		memorydb_describeReservedNodesCmd.Flags().String("reservation-id", "", "The reserved node identifier filter value.")
		memorydb_describeReservedNodesCmd.Flags().String("reserved-nodes-offering-id", "", "The offering identifier filter value.")
	})
	memorydbCmd.AddCommand(memorydb_describeReservedNodesCmd)
}
