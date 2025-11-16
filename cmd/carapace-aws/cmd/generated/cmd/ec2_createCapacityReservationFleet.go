package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createCapacityReservationFleetCmd = &cobra.Command{
	Use:   "create-capacity-reservation-fleet",
	Short: "Creates a Capacity Reservation Fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createCapacityReservationFleetCmd).Standalone()

	ec2_createCapacityReservationFleetCmd.Flags().String("allocation-strategy", "", "The strategy used by the Capacity Reservation Fleet to determine which of the specified instance types to use.")
	ec2_createCapacityReservationFleetCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createCapacityReservationFleetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createCapacityReservationFleetCmd.Flags().String("end-date", "", "The date and time at which the Capacity Reservation Fleet expires.")
	ec2_createCapacityReservationFleetCmd.Flags().String("instance-match-criteria", "", "Indicates the type of instance launches that the Capacity Reservation Fleet accepts.")
	ec2_createCapacityReservationFleetCmd.Flags().String("instance-type-specifications", "", "Information about the instance types for which to reserve the capacity.")
	ec2_createCapacityReservationFleetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createCapacityReservationFleetCmd.Flags().String("tag-specifications", "", "The tags to assign to the Capacity Reservation Fleet.")
	ec2_createCapacityReservationFleetCmd.Flags().String("tenancy", "", "Indicates the tenancy of the Capacity Reservation Fleet.")
	ec2_createCapacityReservationFleetCmd.Flags().String("total-target-capacity", "", "The total number of capacity units to be reserved by the Capacity Reservation Fleet.")
	ec2_createCapacityReservationFleetCmd.MarkFlagRequired("instance-type-specifications")
	ec2_createCapacityReservationFleetCmd.Flag("no-dry-run").Hidden = true
	ec2_createCapacityReservationFleetCmd.MarkFlagRequired("total-target-capacity")
	ec2Cmd.AddCommand(ec2_createCapacityReservationFleetCmd)
}
