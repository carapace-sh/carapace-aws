package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_batchCreateVehicleCmd = &cobra.Command{
	Use:   "batch-create-vehicle",
	Short: "Creates a group, or batch, of vehicles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_batchCreateVehicleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_batchCreateVehicleCmd).Standalone()

		iotfleetwise_batchCreateVehicleCmd.Flags().String("vehicles", "", "A list of information about each vehicle to create.")
		iotfleetwise_batchCreateVehicleCmd.MarkFlagRequired("vehicles")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_batchCreateVehicleCmd)
}
