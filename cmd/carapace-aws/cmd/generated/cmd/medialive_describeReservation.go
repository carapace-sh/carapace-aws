package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeReservationCmd = &cobra.Command{
	Use:   "describe-reservation",
	Short: "Get details for a reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeReservationCmd).Standalone()

		medialive_describeReservationCmd.Flags().String("reservation-id", "", "Unique reservation ID, e.g. '1234567'")
		medialive_describeReservationCmd.MarkFlagRequired("reservation-id")
	})
	medialiveCmd.AddCommand(medialive_describeReservationCmd)
}
