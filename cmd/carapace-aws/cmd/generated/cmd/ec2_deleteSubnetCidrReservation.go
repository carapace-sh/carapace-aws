package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteSubnetCidrReservationCmd = &cobra.Command{
	Use:   "delete-subnet-cidr-reservation",
	Short: "Deletes a subnet CIDR reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteSubnetCidrReservationCmd).Standalone()

	ec2_deleteSubnetCidrReservationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteSubnetCidrReservationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteSubnetCidrReservationCmd.Flags().String("subnet-cidr-reservation-id", "", "The ID of the subnet CIDR reservation.")
	ec2_deleteSubnetCidrReservationCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteSubnetCidrReservationCmd.MarkFlagRequired("subnet-cidr-reservation-id")
	ec2Cmd.AddCommand(ec2_deleteSubnetCidrReservationCmd)
}
