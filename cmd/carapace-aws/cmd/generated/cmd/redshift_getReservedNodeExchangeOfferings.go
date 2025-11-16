package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_getReservedNodeExchangeOfferingsCmd = &cobra.Command{
	Use:   "get-reserved-node-exchange-offerings",
	Short: "Returns an array of DC2 ReservedNodeOfferings that matches the payment type, term, and usage price of the given DC1 reserved node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_getReservedNodeExchangeOfferingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_getReservedNodeExchangeOfferingsCmd).Standalone()

		redshift_getReservedNodeExchangeOfferingsCmd.Flags().String("marker", "", "A value that indicates the starting point for the next set of ReservedNodeOfferings.")
		redshift_getReservedNodeExchangeOfferingsCmd.Flags().String("max-records", "", "An integer setting the maximum number of ReservedNodeOfferings to retrieve.")
		redshift_getReservedNodeExchangeOfferingsCmd.Flags().String("reserved-node-id", "", "A string representing the node identifier for the DC1 Reserved Node to be exchanged.")
		redshift_getReservedNodeExchangeOfferingsCmd.MarkFlagRequired("reserved-node-id")
	})
	redshiftCmd.AddCommand(redshift_getReservedNodeExchangeOfferingsCmd)
}
