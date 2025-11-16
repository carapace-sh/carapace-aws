package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createCapacityReservationBySplittingCmd = &cobra.Command{
	Use:   "create-capacity-reservation-by-splitting",
	Short: "Create a new Capacity Reservation by splitting the capacity of the source Capacity Reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createCapacityReservationBySplittingCmd).Standalone()

	ec2_createCapacityReservationBySplittingCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createCapacityReservationBySplittingCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createCapacityReservationBySplittingCmd.Flags().String("instance-count", "", "The number of instances to split from the source Capacity Reservation.")
	ec2_createCapacityReservationBySplittingCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createCapacityReservationBySplittingCmd.Flags().String("source-capacity-reservation-id", "", "The ID of the Capacity Reservation from which you want to split the capacity.")
	ec2_createCapacityReservationBySplittingCmd.Flags().String("tag-specifications", "", "The tags to apply to the new Capacity Reservation.")
	ec2_createCapacityReservationBySplittingCmd.MarkFlagRequired("instance-count")
	ec2_createCapacityReservationBySplittingCmd.Flag("no-dry-run").Hidden = true
	ec2_createCapacityReservationBySplittingCmd.MarkFlagRequired("source-capacity-reservation-id")
	ec2Cmd.AddCommand(ec2_createCapacityReservationBySplittingCmd)
}
