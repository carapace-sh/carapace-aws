package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateCapacityReservationBillingOwnerCmd = &cobra.Command{
	Use:   "disassociate-capacity-reservation-billing-owner",
	Short: "Cancels a pending request to assign billing of the unused capacity of a Capacity Reservation to a consumer account, or revokes a request that has already been accepted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateCapacityReservationBillingOwnerCmd).Standalone()

	ec2_disassociateCapacityReservationBillingOwnerCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity Reservation.")
	ec2_disassociateCapacityReservationBillingOwnerCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateCapacityReservationBillingOwnerCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateCapacityReservationBillingOwnerCmd.Flags().String("unused-reservation-billing-owner-id", "", "The ID of the consumer account to which the request was sent.")
	ec2_disassociateCapacityReservationBillingOwnerCmd.MarkFlagRequired("capacity-reservation-id")
	ec2_disassociateCapacityReservationBillingOwnerCmd.Flag("no-dry-run").Hidden = true
	ec2_disassociateCapacityReservationBillingOwnerCmd.MarkFlagRequired("unused-reservation-billing-owner-id")
	ec2Cmd.AddCommand(ec2_disassociateCapacityReservationBillingOwnerCmd)
}
