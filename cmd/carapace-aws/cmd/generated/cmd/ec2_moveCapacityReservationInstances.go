package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_moveCapacityReservationInstancesCmd = &cobra.Command{
	Use:   "move-capacity-reservation-instances",
	Short: "Move available capacity from a source Capacity Reservation to a destination Capacity Reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_moveCapacityReservationInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_moveCapacityReservationInstancesCmd).Standalone()

		ec2_moveCapacityReservationInstancesCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_moveCapacityReservationInstancesCmd.Flags().String("destination-capacity-reservation-id", "", "The ID of the Capacity Reservation that you want to move capacity into.")
		ec2_moveCapacityReservationInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_moveCapacityReservationInstancesCmd.Flags().String("instance-count", "", "The number of instances that you want to move from the source Capacity Reservation.")
		ec2_moveCapacityReservationInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_moveCapacityReservationInstancesCmd.Flags().String("source-capacity-reservation-id", "", "The ID of the Capacity Reservation from which you want to move capacity.")
		ec2_moveCapacityReservationInstancesCmd.MarkFlagRequired("destination-capacity-reservation-id")
		ec2_moveCapacityReservationInstancesCmd.MarkFlagRequired("instance-count")
		ec2_moveCapacityReservationInstancesCmd.Flag("no-dry-run").Hidden = true
		ec2_moveCapacityReservationInstancesCmd.MarkFlagRequired("source-capacity-reservation-id")
	})
	ec2Cmd.AddCommand(ec2_moveCapacityReservationInstancesCmd)
}
