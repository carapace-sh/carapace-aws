package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelCapacityReservationFleetsCmd = &cobra.Command{
	Use:   "cancel-capacity-reservation-fleets",
	Short: "Cancels one or more Capacity Reservation Fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelCapacityReservationFleetsCmd).Standalone()

	ec2_cancelCapacityReservationFleetsCmd.Flags().String("capacity-reservation-fleet-ids", "", "The IDs of the Capacity Reservation Fleets to cancel.")
	ec2_cancelCapacityReservationFleetsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_cancelCapacityReservationFleetsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_cancelCapacityReservationFleetsCmd.MarkFlagRequired("capacity-reservation-fleet-ids")
	ec2_cancelCapacityReservationFleetsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_cancelCapacityReservationFleetsCmd)
}
