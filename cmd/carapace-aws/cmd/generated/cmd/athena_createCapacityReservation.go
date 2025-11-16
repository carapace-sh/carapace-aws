package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_createCapacityReservationCmd = &cobra.Command{
	Use:   "create-capacity-reservation",
	Short: "Creates a capacity reservation with the specified name and number of requested data processing units.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_createCapacityReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_createCapacityReservationCmd).Standalone()

		athena_createCapacityReservationCmd.Flags().String("name", "", "The name of the capacity reservation to create.")
		athena_createCapacityReservationCmd.Flags().String("tags", "", "The tags for the capacity reservation.")
		athena_createCapacityReservationCmd.Flags().String("target-dpus", "", "The number of requested data processing units.")
		athena_createCapacityReservationCmd.MarkFlagRequired("name")
		athena_createCapacityReservationCmd.MarkFlagRequired("target-dpus")
	})
	athenaCmd.AddCommand(athena_createCapacityReservationCmd)
}
