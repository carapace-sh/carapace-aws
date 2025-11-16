package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_disassociateVehicleFleetCmd = &cobra.Command{
	Use:   "disassociate-vehicle-fleet",
	Short: "Removes, or disassociates, a vehicle from a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_disassociateVehicleFleetCmd).Standalone()

	iotfleetwise_disassociateVehicleFleetCmd.Flags().String("fleet-id", "", "The unique ID of a fleet.")
	iotfleetwise_disassociateVehicleFleetCmd.Flags().String("vehicle-name", "", "The unique ID of the vehicle to disassociate from the fleet.")
	iotfleetwise_disassociateVehicleFleetCmd.MarkFlagRequired("fleet-id")
	iotfleetwise_disassociateVehicleFleetCmd.MarkFlagRequired("vehicle-name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_disassociateVehicleFleetCmd)
}
