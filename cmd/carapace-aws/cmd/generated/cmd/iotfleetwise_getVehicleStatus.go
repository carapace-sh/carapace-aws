package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getVehicleStatusCmd = &cobra.Command{
	Use:   "get-vehicle-status",
	Short: "Retrieves information about the status of campaigns, decoder manifests, or state templates associated with a vehicle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getVehicleStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_getVehicleStatusCmd).Standalone()

		iotfleetwise_getVehicleStatusCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_getVehicleStatusCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
		iotfleetwise_getVehicleStatusCmd.Flags().String("vehicle-name", "", "The ID of the vehicle to retrieve information about.")
		iotfleetwise_getVehicleStatusCmd.MarkFlagRequired("vehicle-name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_getVehicleStatusCmd)
}
