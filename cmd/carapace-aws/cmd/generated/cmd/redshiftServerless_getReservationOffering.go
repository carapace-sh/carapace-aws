package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getReservationOfferingCmd = &cobra.Command{
	Use:   "get-reservation-offering",
	Short: "Returns the reservation offering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getReservationOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_getReservationOfferingCmd).Standalone()

		redshiftServerless_getReservationOfferingCmd.Flags().String("offering-id", "", "The identifier for the offering..")
		redshiftServerless_getReservationOfferingCmd.MarkFlagRequired("offering-id")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_getReservationOfferingCmd)
}
