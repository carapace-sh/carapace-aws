package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyCapacityReservationCmd = &cobra.Command{
	Use:   "modify-capacity-reservation",
	Short: "Modifies a Capacity Reservation's capacity, instance eligibility, and the conditions under which it is to be released.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyCapacityReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyCapacityReservationCmd).Standalone()

		ec2_modifyCapacityReservationCmd.Flags().Bool("accept", false, "Reserved.")
		ec2_modifyCapacityReservationCmd.Flags().String("additional-info", "", "Reserved for future use.")
		ec2_modifyCapacityReservationCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity Reservation.")
		ec2_modifyCapacityReservationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyCapacityReservationCmd.Flags().String("end-date", "", "The date and time at which the Capacity Reservation expires.")
		ec2_modifyCapacityReservationCmd.Flags().String("end-date-type", "", "Indicates the way in which the Capacity Reservation ends.")
		ec2_modifyCapacityReservationCmd.Flags().String("instance-count", "", "The number of instances for which to reserve capacity.")
		ec2_modifyCapacityReservationCmd.Flags().String("instance-match-criteria", "", "The matching criteria (instance eligibility) that you want to use in the modified Capacity Reservation.")
		ec2_modifyCapacityReservationCmd.Flags().Bool("no-accept", false, "Reserved.")
		ec2_modifyCapacityReservationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyCapacityReservationCmd.MarkFlagRequired("capacity-reservation-id")
		ec2_modifyCapacityReservationCmd.Flag("no-accept").Hidden = true
		ec2_modifyCapacityReservationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyCapacityReservationCmd)
}
