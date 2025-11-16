package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeReservedDbinstancesOfferingsCmd = &cobra.Command{
	Use:   "describe-reserved-dbinstances-offerings",
	Short: "Lists available reserved DB instance offerings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeReservedDbinstancesOfferingsCmd).Standalone()

	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("dbinstance-class", "", "The DB instance class filter value.")
	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("duration", "", "Duration filter value, specified in years or seconds.")
	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("multi-az", "", "Specifies whether to show only those reservations that support Multi-AZ.")
	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("offering-type", "", "The offering type filter value.")
	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("product-description", "", "Product description filter value.")
	rds_describeReservedDbinstancesOfferingsCmd.Flags().String("reserved-dbinstances-offering-id", "", "The offering identifier filter value.")
	rdsCmd.AddCommand(rds_describeReservedDbinstancesOfferingsCmd)
}
