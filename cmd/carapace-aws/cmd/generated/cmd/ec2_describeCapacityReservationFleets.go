package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityReservationFleetsCmd = &cobra.Command{
	Use:   "describe-capacity-reservation-fleets",
	Short: "Describes one or more Capacity Reservation Fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityReservationFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCapacityReservationFleetsCmd).Standalone()

		ec2_describeCapacityReservationFleetsCmd.Flags().String("capacity-reservation-fleet-ids", "", "The IDs of the Capacity Reservation Fleets to describe.")
		ec2_describeCapacityReservationFleetsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityReservationFleetsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeCapacityReservationFleetsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeCapacityReservationFleetsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		ec2_describeCapacityReservationFleetsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityReservationFleetsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCapacityReservationFleetsCmd)
}
