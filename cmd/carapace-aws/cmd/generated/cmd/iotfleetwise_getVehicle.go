package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getVehicleCmd = &cobra.Command{
	Use:   "get-vehicle",
	Short: "Retrieves information about a vehicle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getVehicleCmd).Standalone()

	iotfleetwise_getVehicleCmd.Flags().String("vehicle-name", "", "The ID of the vehicle to retrieve information about.")
	iotfleetwise_getVehicleCmd.MarkFlagRequired("vehicle-name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_getVehicleCmd)
}
