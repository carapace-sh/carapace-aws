package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_deleteCapacityReservationCmd = &cobra.Command{
	Use:   "delete-capacity-reservation",
	Short: "Deletes a cancelled capacity reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_deleteCapacityReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_deleteCapacityReservationCmd).Standalone()

		athena_deleteCapacityReservationCmd.Flags().String("name", "", "The name of the capacity reservation to delete.")
		athena_deleteCapacityReservationCmd.MarkFlagRequired("name")
	})
	athenaCmd.AddCommand(athena_deleteCapacityReservationCmd)
}
