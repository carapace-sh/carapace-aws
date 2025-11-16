package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeReservedNodeExchangeStatusCmd = &cobra.Command{
	Use:   "describe-reserved-node-exchange-status",
	Short: "Returns exchange status details and associated metadata for a reserved-node exchange.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeReservedNodeExchangeStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeReservedNodeExchangeStatusCmd).Standalone()

		redshift_describeReservedNodeExchangeStatusCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeReservedNodeExchangeStatus` request.")
		redshift_describeReservedNodeExchangeStatusCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeReservedNodeExchangeStatusCmd.Flags().String("reserved-node-exchange-request-id", "", "The identifier of the reserved-node exchange request.")
		redshift_describeReservedNodeExchangeStatusCmd.Flags().String("reserved-node-id", "", "The identifier of the source reserved node in a reserved-node exchange request.")
	})
	redshiftCmd.AddCommand(redshift_describeReservedNodeExchangeStatusCmd)
}
