package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_rejectCapacityReservationBillingOwnershipCmd = &cobra.Command{
	Use:   "reject-capacity-reservation-billing-ownership",
	Short: "Rejects a request to assign billing of the available capacity of a shared Capacity Reservation to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_rejectCapacityReservationBillingOwnershipCmd).Standalone()

	ec2_rejectCapacityReservationBillingOwnershipCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity Reservation for which to reject the request.")
	ec2_rejectCapacityReservationBillingOwnershipCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_rejectCapacityReservationBillingOwnershipCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_rejectCapacityReservationBillingOwnershipCmd.MarkFlagRequired("capacity-reservation-id")
	ec2_rejectCapacityReservationBillingOwnershipCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_rejectCapacityReservationBillingOwnershipCmd)
}
