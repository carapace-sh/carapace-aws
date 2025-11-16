package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getCapacityReservationUsageCmd = &cobra.Command{
	Use:   "get-capacity-reservation-usage",
	Short: "Gets usage information about a Capacity Reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getCapacityReservationUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getCapacityReservationUsageCmd).Standalone()

		ec2_getCapacityReservationUsageCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity Reservation.")
		ec2_getCapacityReservationUsageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getCapacityReservationUsageCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_getCapacityReservationUsageCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		ec2_getCapacityReservationUsageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getCapacityReservationUsageCmd.MarkFlagRequired("capacity-reservation-id")
		ec2_getCapacityReservationUsageCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getCapacityReservationUsageCmd)
}
