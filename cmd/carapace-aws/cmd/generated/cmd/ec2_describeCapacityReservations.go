package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityReservationsCmd = &cobra.Command{
	Use:   "describe-capacity-reservations",
	Short: "Describes one or more of your Capacity Reservations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityReservationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCapacityReservationsCmd).Standalone()

		ec2_describeCapacityReservationsCmd.Flags().String("capacity-reservation-ids", "", "The ID of the Capacity Reservation.")
		ec2_describeCapacityReservationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityReservationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeCapacityReservationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeCapacityReservationsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		ec2_describeCapacityReservationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityReservationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCapacityReservationsCmd)
}
