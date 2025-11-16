package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateReservationCmd = &cobra.Command{
	Use:   "update-reservation",
	Short: "Update reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateReservationCmd).Standalone()

	medialive_updateReservationCmd.Flags().String("name", "", "Name of the reservation")
	medialive_updateReservationCmd.Flags().String("renewal-settings", "", "Renewal settings for the reservation")
	medialive_updateReservationCmd.Flags().String("reservation-id", "", "Unique reservation ID, e.g. '1234567'")
	medialive_updateReservationCmd.MarkFlagRequired("reservation-id")
	medialiveCmd.AddCommand(medialive_updateReservationCmd)
}
