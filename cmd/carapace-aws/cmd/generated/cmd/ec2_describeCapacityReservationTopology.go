package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityReservationTopologyCmd = &cobra.Command{
	Use:   "describe-capacity-reservation-topology",
	Short: "Describes a tree-based hierarchy that represents the physical host placement of your pending or active Capacity Reservations within an Availability Zone or Local Zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityReservationTopologyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCapacityReservationTopologyCmd).Standalone()

		ec2_describeCapacityReservationTopologyCmd.Flags().String("capacity-reservation-ids", "", "The Capacity Reservation IDs.")
		ec2_describeCapacityReservationTopologyCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_describeCapacityReservationTopologyCmd.Flags().String("filters", "", "The filters.")
		ec2_describeCapacityReservationTopologyCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeCapacityReservationTopologyCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeCapacityReservationTopologyCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_describeCapacityReservationTopologyCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCapacityReservationTopologyCmd)
}
