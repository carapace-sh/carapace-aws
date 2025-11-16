package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createSubnetCidrReservationCmd = &cobra.Command{
	Use:   "create-subnet-cidr-reservation",
	Short: "Creates a subnet CIDR reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createSubnetCidrReservationCmd).Standalone()

	ec2_createSubnetCidrReservationCmd.Flags().String("cidr", "", "The IPv4 or IPV6 CIDR range to reserve.")
	ec2_createSubnetCidrReservationCmd.Flags().String("description", "", "The description to assign to the subnet CIDR reservation.")
	ec2_createSubnetCidrReservationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createSubnetCidrReservationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createSubnetCidrReservationCmd.Flags().String("reservation-type", "", "The type of reservation.")
	ec2_createSubnetCidrReservationCmd.Flags().String("subnet-id", "", "The ID of the subnet.")
	ec2_createSubnetCidrReservationCmd.Flags().String("tag-specifications", "", "The tags to assign to the subnet CIDR reservation.")
	ec2_createSubnetCidrReservationCmd.MarkFlagRequired("cidr")
	ec2_createSubnetCidrReservationCmd.Flag("no-dry-run").Hidden = true
	ec2_createSubnetCidrReservationCmd.MarkFlagRequired("reservation-type")
	ec2_createSubnetCidrReservationCmd.MarkFlagRequired("subnet-id")
	ec2Cmd.AddCommand(ec2_createSubnetCidrReservationCmd)
}
