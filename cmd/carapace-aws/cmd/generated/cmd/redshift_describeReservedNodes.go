package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeReservedNodesCmd = &cobra.Command{
	Use:   "describe-reserved-nodes",
	Short: "Returns the descriptions of the reserved nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeReservedNodesCmd).Standalone()

	redshift_describeReservedNodesCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeReservedNodesCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeReservedNodesCmd.Flags().String("reserved-node-id", "", "Identifier for the node reservation.")
	redshiftCmd.AddCommand(redshift_describeReservedNodesCmd)
}
