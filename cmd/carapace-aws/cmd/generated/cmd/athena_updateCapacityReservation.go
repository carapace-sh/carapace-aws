package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_updateCapacityReservationCmd = &cobra.Command{
	Use:   "update-capacity-reservation",
	Short: "Updates the number of requested data processing units for the capacity reservation with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_updateCapacityReservationCmd).Standalone()

	athena_updateCapacityReservationCmd.Flags().String("name", "", "The name of the capacity reservation.")
	athena_updateCapacityReservationCmd.Flags().String("target-dpus", "", "The new number of requested data processing units.")
	athena_updateCapacityReservationCmd.MarkFlagRequired("name")
	athena_updateCapacityReservationCmd.MarkFlagRequired("target-dpus")
	athenaCmd.AddCommand(athena_updateCapacityReservationCmd)
}
