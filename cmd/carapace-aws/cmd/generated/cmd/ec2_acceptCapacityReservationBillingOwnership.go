package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_acceptCapacityReservationBillingOwnershipCmd = &cobra.Command{
	Use:   "accept-capacity-reservation-billing-ownership",
	Short: "Accepts a request to assign billing of the available capacity of a shared Capacity Reservation to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_acceptCapacityReservationBillingOwnershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_acceptCapacityReservationBillingOwnershipCmd).Standalone()

		ec2_acceptCapacityReservationBillingOwnershipCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity Reservation for which to accept the request.")
		ec2_acceptCapacityReservationBillingOwnershipCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_acceptCapacityReservationBillingOwnershipCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_acceptCapacityReservationBillingOwnershipCmd.MarkFlagRequired("capacity-reservation-id")
		ec2_acceptCapacityReservationBillingOwnershipCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_acceptCapacityReservationBillingOwnershipCmd)
}
