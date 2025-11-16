package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listReservationOfferingsCmd = &cobra.Command{
	Use:   "list-reservation-offerings",
	Short: "Returns the current reservation offerings in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listReservationOfferingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_listReservationOfferingsCmd).Standalone()

		redshiftServerless_listReservationOfferingsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		redshiftServerless_listReservationOfferingsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_listReservationOfferingsCmd)
}
