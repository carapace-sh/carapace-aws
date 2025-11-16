package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getCapacityReservationCmd = &cobra.Command{
	Use:   "get-capacity-reservation",
	Short: "Returns information about the capacity reservation with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getCapacityReservationCmd).Standalone()

	athena_getCapacityReservationCmd.Flags().String("name", "", "The name of the capacity reservation.")
	athena_getCapacityReservationCmd.MarkFlagRequired("name")
	athenaCmd.AddCommand(athena_getCapacityReservationCmd)
}
