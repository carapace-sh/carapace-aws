package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceCapacityReservationAttributesCmd = &cobra.Command{
	Use:   "modify-instance-capacity-reservation-attributes",
	Short: "Modifies the Capacity Reservation settings for a stopped instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceCapacityReservationAttributesCmd).Standalone()

	ec2_modifyInstanceCapacityReservationAttributesCmd.Flags().String("capacity-reservation-specification", "", "Information about the Capacity Reservation targeting option.")
	ec2_modifyInstanceCapacityReservationAttributesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyInstanceCapacityReservationAttributesCmd.Flags().String("instance-id", "", "The ID of the instance to be modified.")
	ec2_modifyInstanceCapacityReservationAttributesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyInstanceCapacityReservationAttributesCmd.MarkFlagRequired("capacity-reservation-specification")
	ec2_modifyInstanceCapacityReservationAttributesCmd.MarkFlagRequired("instance-id")
	ec2_modifyInstanceCapacityReservationAttributesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyInstanceCapacityReservationAttributesCmd)
}
