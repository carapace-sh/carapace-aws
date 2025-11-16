package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteReservationCmd = &cobra.Command{
	Use:   "delete-reservation",
	Short: "Delete an expired reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_deleteReservationCmd).Standalone()

		medialive_deleteReservationCmd.Flags().String("reservation-id", "", "Unique reservation ID, e.g. '1234567'")
		medialive_deleteReservationCmd.MarkFlagRequired("reservation-id")
	})
	medialiveCmd.AddCommand(medialive_deleteReservationCmd)
}
