package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_purchaseHostReservationCmd = &cobra.Command{
	Use:   "purchase-host-reservation",
	Short: "Purchase a reservation with configurations that match those of your Dedicated Host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_purchaseHostReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_purchaseHostReservationCmd).Standalone()

		ec2_purchaseHostReservationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_purchaseHostReservationCmd.Flags().String("currency-code", "", "The currency in which the `totalUpfrontPrice`, `LimitPrice`, and `totalHourlyPrice` amounts are specified.")
		ec2_purchaseHostReservationCmd.Flags().String("host-id-set", "", "The IDs of the Dedicated Hosts with which the reservation will be associated.")
		ec2_purchaseHostReservationCmd.Flags().String("limit-price", "", "The specified limit is checked against the total upfront cost of the reservation (calculated as the offering's upfront cost multiplied by the host count).")
		ec2_purchaseHostReservationCmd.Flags().String("offering-id", "", "The ID of the offering.")
		ec2_purchaseHostReservationCmd.Flags().String("tag-specifications", "", "The tags to apply to the Dedicated Host Reservation during purchase.")
		ec2_purchaseHostReservationCmd.MarkFlagRequired("host-id-set")
		ec2_purchaseHostReservationCmd.MarkFlagRequired("offering-id")
	})
	ec2Cmd.AddCommand(ec2_purchaseHostReservationCmd)
}
