package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_deleteVehicleCmd = &cobra.Command{
	Use:   "delete-vehicle",
	Short: "Deletes a vehicle and removes it from any campaigns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_deleteVehicleCmd).Standalone()

	iotfleetwise_deleteVehicleCmd.Flags().String("vehicle-name", "", "The ID of the vehicle to delete.")
	iotfleetwise_deleteVehicleCmd.MarkFlagRequired("vehicle-name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_deleteVehicleCmd)
}
