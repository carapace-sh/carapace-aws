package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_associateVehicleFleetCmd = &cobra.Command{
	Use:   "associate-vehicle-fleet",
	Short: "Adds, or associates, a vehicle with a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_associateVehicleFleetCmd).Standalone()

	iotfleetwise_associateVehicleFleetCmd.Flags().String("fleet-id", "", "The ID of a fleet.")
	iotfleetwise_associateVehicleFleetCmd.Flags().String("vehicle-name", "", "The unique ID of the vehicle to associate with the fleet.")
	iotfleetwise_associateVehicleFleetCmd.MarkFlagRequired("fleet-id")
	iotfleetwise_associateVehicleFleetCmd.MarkFlagRequired("vehicle-name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_associateVehicleFleetCmd)
}
