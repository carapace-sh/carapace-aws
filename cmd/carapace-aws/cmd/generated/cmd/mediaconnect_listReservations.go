package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_listReservationsCmd = &cobra.Command{
	Use:   "list-reservations",
	Short: "Displays a list of all reservations that have been purchased by this account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_listReservationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_listReservationsCmd).Standalone()

		mediaconnect_listReservationsCmd.Flags().String("max-results", "", "The maximum number of results to return per API request.")
		mediaconnect_listReservationsCmd.Flags().String("next-token", "", "The token that identifies the batch of results that you want to see.")
	})
	mediaconnectCmd.AddCommand(mediaconnect_listReservationsCmd)
}
