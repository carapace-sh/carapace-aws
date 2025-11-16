package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeHostReservationsCmd = &cobra.Command{
	Use:   "describe-host-reservations",
	Short: "Describes reservations that are associated with Dedicated Hosts in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeHostReservationsCmd).Standalone()

	ec2_describeHostReservationsCmd.Flags().String("filter", "", "The filters.")
	ec2_describeHostReservationsCmd.Flags().String("host-reservation-id-set", "", "The host reservation IDs.")
	ec2_describeHostReservationsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeHostReservationsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	ec2Cmd.AddCommand(ec2_describeHostReservationsCmd)
}
