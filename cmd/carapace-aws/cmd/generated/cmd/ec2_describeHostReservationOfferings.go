package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeHostReservationOfferingsCmd = &cobra.Command{
	Use:   "describe-host-reservation-offerings",
	Short: "Describes the Dedicated Host reservations that are available to purchase.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeHostReservationOfferingsCmd).Standalone()

	ec2_describeHostReservationOfferingsCmd.Flags().String("filter", "", "The filters.")
	ec2_describeHostReservationOfferingsCmd.Flags().String("max-duration", "", "This is the maximum duration of the reservation to purchase, specified in seconds.")
	ec2_describeHostReservationOfferingsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeHostReservationOfferingsCmd.Flags().String("min-duration", "", "This is the minimum duration of the reservation you'd like to purchase, specified in seconds.")
	ec2_describeHostReservationOfferingsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	ec2_describeHostReservationOfferingsCmd.Flags().String("offering-id", "", "The ID of the reservation offering.")
	ec2Cmd.AddCommand(ec2_describeHostReservationOfferingsCmd)
}
