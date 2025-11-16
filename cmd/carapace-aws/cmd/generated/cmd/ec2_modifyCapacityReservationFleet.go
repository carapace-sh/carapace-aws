package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyCapacityReservationFleetCmd = &cobra.Command{
	Use:   "modify-capacity-reservation-fleet",
	Short: "Modifies a Capacity Reservation Fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyCapacityReservationFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyCapacityReservationFleetCmd).Standalone()

		ec2_modifyCapacityReservationFleetCmd.Flags().String("capacity-reservation-fleet-id", "", "The ID of the Capacity Reservation Fleet to modify.")
		ec2_modifyCapacityReservationFleetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyCapacityReservationFleetCmd.Flags().String("end-date", "", "The date and time at which the Capacity Reservation Fleet expires.")
		ec2_modifyCapacityReservationFleetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyCapacityReservationFleetCmd.Flags().Bool("no-remove-end-date", false, "Indicates whether to remove the end date from the Capacity Reservation Fleet.")
		ec2_modifyCapacityReservationFleetCmd.Flags().Bool("remove-end-date", false, "Indicates whether to remove the end date from the Capacity Reservation Fleet.")
		ec2_modifyCapacityReservationFleetCmd.Flags().String("total-target-capacity", "", "The total number of capacity units to be reserved by the Capacity Reservation Fleet.")
		ec2_modifyCapacityReservationFleetCmd.MarkFlagRequired("capacity-reservation-fleet-id")
		ec2_modifyCapacityReservationFleetCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyCapacityReservationFleetCmd.Flag("no-remove-end-date").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyCapacityReservationFleetCmd)
}
