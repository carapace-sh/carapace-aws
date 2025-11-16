package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listReservationsCmd = &cobra.Command{
	Use:   "list-reservations",
	Short: "Returns a list of Reservation objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listReservationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_listReservationsCmd).Standalone()

		redshiftServerless_listReservationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		redshiftServerless_listReservationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_listReservationsCmd)
}
