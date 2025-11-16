package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getSubnetCidrReservationsCmd = &cobra.Command{
	Use:   "get-subnet-cidr-reservations",
	Short: "Gets information about the subnet CIDR reservations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getSubnetCidrReservationsCmd).Standalone()

	ec2_getSubnetCidrReservationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getSubnetCidrReservationsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_getSubnetCidrReservationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_getSubnetCidrReservationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getSubnetCidrReservationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getSubnetCidrReservationsCmd.Flags().String("subnet-id", "", "The ID of the subnet.")
	ec2_getSubnetCidrReservationsCmd.Flag("no-dry-run").Hidden = true
	ec2_getSubnetCidrReservationsCmd.MarkFlagRequired("subnet-id")
	ec2Cmd.AddCommand(ec2_getSubnetCidrReservationsCmd)
}
