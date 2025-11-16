package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_cancelCapacityReservationCmd = &cobra.Command{
	Use:   "cancel-capacity-reservation",
	Short: "Cancels the capacity reservation with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_cancelCapacityReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_cancelCapacityReservationCmd).Standalone()

		athena_cancelCapacityReservationCmd.Flags().String("name", "", "The name of the capacity reservation to cancel.")
		athena_cancelCapacityReservationCmd.MarkFlagRequired("name")
	})
	athenaCmd.AddCommand(athena_cancelCapacityReservationCmd)
}
