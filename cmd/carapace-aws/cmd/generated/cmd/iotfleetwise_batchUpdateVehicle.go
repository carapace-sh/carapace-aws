package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_batchUpdateVehicleCmd = &cobra.Command{
	Use:   "batch-update-vehicle",
	Short: "Updates a group, or batch, of vehicles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_batchUpdateVehicleCmd).Standalone()

	iotfleetwise_batchUpdateVehicleCmd.Flags().String("vehicles", "", "A list of information about the vehicles to update.")
	iotfleetwise_batchUpdateVehicleCmd.MarkFlagRequired("vehicles")
	iotfleetwiseCmd.AddCommand(iotfleetwise_batchUpdateVehicleCmd)
}
