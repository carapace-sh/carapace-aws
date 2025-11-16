package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listFleetsForVehicleCmd = &cobra.Command{
	Use:   "list-fleets-for-vehicle",
	Short: "Retrieves a list of IDs for all fleets that the vehicle is associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listFleetsForVehicleCmd).Standalone()

	iotfleetwise_listFleetsForVehicleCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
	iotfleetwise_listFleetsForVehicleCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	iotfleetwise_listFleetsForVehicleCmd.Flags().String("vehicle-name", "", "The ID of the vehicle to retrieve information about.")
	iotfleetwise_listFleetsForVehicleCmd.MarkFlagRequired("vehicle-name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_listFleetsForVehicleCmd)
}
