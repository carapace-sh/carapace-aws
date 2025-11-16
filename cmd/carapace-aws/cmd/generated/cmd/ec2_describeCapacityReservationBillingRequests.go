package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityReservationBillingRequestsCmd = &cobra.Command{
	Use:   "describe-capacity-reservation-billing-requests",
	Short: "Describes a request to assign the billing of the unused capacity of a Capacity Reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityReservationBillingRequestsCmd).Standalone()

	ec2_describeCapacityReservationBillingRequestsCmd.Flags().String("capacity-reservation-ids", "", "The ID of the Capacity Reservation.")
	ec2_describeCapacityReservationBillingRequestsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeCapacityReservationBillingRequestsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeCapacityReservationBillingRequestsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeCapacityReservationBillingRequestsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	ec2_describeCapacityReservationBillingRequestsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeCapacityReservationBillingRequestsCmd.Flags().String("role", "", "Specify one of the following:")
	ec2_describeCapacityReservationBillingRequestsCmd.Flag("no-dry-run").Hidden = true
	ec2_describeCapacityReservationBillingRequestsCmd.MarkFlagRequired("role")
	ec2Cmd.AddCommand(ec2_describeCapacityReservationBillingRequestsCmd)
}
