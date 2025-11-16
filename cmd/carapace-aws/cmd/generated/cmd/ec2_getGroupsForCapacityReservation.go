package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getGroupsForCapacityReservationCmd = &cobra.Command{
	Use:   "get-groups-for-capacity-reservation",
	Short: "Lists the resource groups to which a Capacity Reservation has been added.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getGroupsForCapacityReservationCmd).Standalone()

	ec2_getGroupsForCapacityReservationCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity Reservation.")
	ec2_getGroupsForCapacityReservationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getGroupsForCapacityReservationCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_getGroupsForCapacityReservationCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	ec2_getGroupsForCapacityReservationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getGroupsForCapacityReservationCmd.MarkFlagRequired("capacity-reservation-id")
	ec2_getGroupsForCapacityReservationCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getGroupsForCapacityReservationCmd)
}
