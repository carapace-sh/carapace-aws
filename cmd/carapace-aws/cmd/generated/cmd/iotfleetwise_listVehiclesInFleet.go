package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listVehiclesInFleetCmd = &cobra.Command{
	Use:   "list-vehicles-in-fleet",
	Short: "Retrieves a list of summaries of all vehicles associated with a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listVehiclesInFleetCmd).Standalone()

	iotfleetwise_listVehiclesInFleetCmd.Flags().String("fleet-id", "", "The ID of a fleet.")
	iotfleetwise_listVehiclesInFleetCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
	iotfleetwise_listVehiclesInFleetCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	iotfleetwise_listVehiclesInFleetCmd.MarkFlagRequired("fleet-id")
	iotfleetwiseCmd.AddCommand(iotfleetwise_listVehiclesInFleetCmd)
}
