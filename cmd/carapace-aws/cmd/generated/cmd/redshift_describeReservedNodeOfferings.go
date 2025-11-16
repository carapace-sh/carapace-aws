package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeReservedNodeOfferingsCmd = &cobra.Command{
	Use:   "describe-reserved-node-offerings",
	Short: "Returns a list of the available reserved node offerings by Amazon Redshift with their descriptions including the node type, the fixed and recurring costs of reserving the node and duration the node will be reserved for you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeReservedNodeOfferingsCmd).Standalone()

	redshift_describeReservedNodeOfferingsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeReservedNodeOfferingsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeReservedNodeOfferingsCmd.Flags().String("reserved-node-offering-id", "", "The unique identifier for the offering.")
	redshiftCmd.AddCommand(redshift_describeReservedNodeOfferingsCmd)
}
