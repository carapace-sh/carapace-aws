package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelCapacityReservationCmd = &cobra.Command{
	Use:   "cancel-capacity-reservation",
	Short: "Cancels the specified Capacity Reservation, releases the reserved capacity, and changes the Capacity Reservation's state to `cancelled`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelCapacityReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_cancelCapacityReservationCmd).Standalone()

		ec2_cancelCapacityReservationCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity Reservation to be cancelled.")
		ec2_cancelCapacityReservationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelCapacityReservationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelCapacityReservationCmd.MarkFlagRequired("capacity-reservation-id")
		ec2_cancelCapacityReservationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_cancelCapacityReservationCmd)
}
