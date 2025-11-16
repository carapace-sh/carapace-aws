package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateCapacityReservationBillingOwnerCmd = &cobra.Command{
	Use:   "associate-capacity-reservation-billing-owner",
	Short: "Initiates a request to assign billing of the unused capacity of a shared Capacity Reservation to a consumer account that is consolidated under the same Amazon Web Services organizations payer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateCapacityReservationBillingOwnerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_associateCapacityReservationBillingOwnerCmd).Standalone()

		ec2_associateCapacityReservationBillingOwnerCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity Reservation.")
		ec2_associateCapacityReservationBillingOwnerCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateCapacityReservationBillingOwnerCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateCapacityReservationBillingOwnerCmd.Flags().String("unused-reservation-billing-owner-id", "", "The ID of the consumer account to which to assign billing.")
		ec2_associateCapacityReservationBillingOwnerCmd.MarkFlagRequired("capacity-reservation-id")
		ec2_associateCapacityReservationBillingOwnerCmd.Flag("no-dry-run").Hidden = true
		ec2_associateCapacityReservationBillingOwnerCmd.MarkFlagRequired("unused-reservation-billing-owner-id")
	})
	ec2Cmd.AddCommand(ec2_associateCapacityReservationBillingOwnerCmd)
}
