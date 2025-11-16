package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createCapacityReservationCmd = &cobra.Command{
	Use:   "create-capacity-reservation",
	Short: "Creates a new Capacity Reservation with the specified attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createCapacityReservationCmd).Standalone()

	ec2_createCapacityReservationCmd.Flags().String("availability-zone", "", "The Availability Zone in which to create the Capacity Reservation.")
	ec2_createCapacityReservationCmd.Flags().String("availability-zone-id", "", "The ID of the Availability Zone in which to create the Capacity Reservation.")
	ec2_createCapacityReservationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createCapacityReservationCmd.Flags().String("commitment-duration", "", "Required for future-dated Capacity Reservations only.")
	ec2_createCapacityReservationCmd.Flags().String("delivery-preference", "", "Required for future-dated Capacity Reservations only.")
	ec2_createCapacityReservationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createCapacityReservationCmd.Flags().Bool("ebs-optimized", false, "Indicates whether the Capacity Reservation supports EBS-optimized instances.")
	ec2_createCapacityReservationCmd.Flags().String("end-date", "", "The date and time at which the Capacity Reservation expires.")
	ec2_createCapacityReservationCmd.Flags().String("end-date-type", "", "Indicates the way in which the Capacity Reservation ends.")
	ec2_createCapacityReservationCmd.Flags().Bool("ephemeral-storage", false, "*Deprecated.*")
	ec2_createCapacityReservationCmd.Flags().String("instance-count", "", "The number of instances for which to reserve capacity.")
	ec2_createCapacityReservationCmd.Flags().String("instance-match-criteria", "", "Indicates the type of instance launches that the Capacity Reservation accepts.")
	ec2_createCapacityReservationCmd.Flags().String("instance-platform", "", "The type of operating system for which to reserve capacity.")
	ec2_createCapacityReservationCmd.Flags().String("instance-type", "", "The instance type for which to reserve capacity.")
	ec2_createCapacityReservationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createCapacityReservationCmd.Flags().Bool("no-ebs-optimized", false, "Indicates whether the Capacity Reservation supports EBS-optimized instances.")
	ec2_createCapacityReservationCmd.Flags().Bool("no-ephemeral-storage", false, "*Deprecated.*")
	ec2_createCapacityReservationCmd.Flags().String("outpost-arn", "", "Not supported for future-dated Capacity Reservations.")
	ec2_createCapacityReservationCmd.Flags().String("placement-group-arn", "", "Not supported for future-dated Capacity Reservations.")
	ec2_createCapacityReservationCmd.Flags().String("start-date", "", "Required for future-dated Capacity Reservations only.")
	ec2_createCapacityReservationCmd.Flags().String("tag-specifications", "", "The tags to apply to the Capacity Reservation during launch.")
	ec2_createCapacityReservationCmd.Flags().String("tenancy", "", "Indicates the tenancy of the Capacity Reservation.")
	ec2_createCapacityReservationCmd.MarkFlagRequired("instance-count")
	ec2_createCapacityReservationCmd.MarkFlagRequired("instance-platform")
	ec2_createCapacityReservationCmd.MarkFlagRequired("instance-type")
	ec2_createCapacityReservationCmd.Flag("no-dry-run").Hidden = true
	ec2_createCapacityReservationCmd.Flag("no-ebs-optimized").Hidden = true
	ec2_createCapacityReservationCmd.Flag("no-ephemeral-storage").Hidden = true
	ec2Cmd.AddCommand(ec2_createCapacityReservationCmd)
}
