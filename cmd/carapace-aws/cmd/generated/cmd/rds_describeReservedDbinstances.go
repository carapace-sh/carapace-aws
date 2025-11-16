package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeReservedDbinstancesCmd = &cobra.Command{
	Use:   "describe-reserved-dbinstances",
	Short: "Returns information about reserved DB instances for this account, or about a specified reserved DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeReservedDbinstancesCmd).Standalone()

	rds_describeReservedDbinstancesCmd.Flags().String("dbinstance-class", "", "The DB instance class filter value.")
	rds_describeReservedDbinstancesCmd.Flags().String("duration", "", "The duration filter value, specified in years or seconds.")
	rds_describeReservedDbinstancesCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
	rds_describeReservedDbinstancesCmd.Flags().String("lease-id", "", "The lease identifier filter value.")
	rds_describeReservedDbinstancesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	rds_describeReservedDbinstancesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeReservedDbinstancesCmd.Flags().String("multi-az", "", "Specifies whether to show only those reservations that support Multi-AZ.")
	rds_describeReservedDbinstancesCmd.Flags().String("offering-type", "", "The offering type filter value.")
	rds_describeReservedDbinstancesCmd.Flags().String("product-description", "", "The product description filter value.")
	rds_describeReservedDbinstancesCmd.Flags().String("reserved-dbinstance-id", "", "The reserved DB instance identifier filter value.")
	rds_describeReservedDbinstancesCmd.Flags().String("reserved-dbinstances-offering-id", "", "The offering identifier filter value.")
	rdsCmd.AddCommand(rds_describeReservedDbinstancesCmd)
}
