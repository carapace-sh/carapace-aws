package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_updateVehicleCmd = &cobra.Command{
	Use:   "update-vehicle",
	Short: "Updates a vehicle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_updateVehicleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_updateVehicleCmd).Standalone()

		iotfleetwise_updateVehicleCmd.Flags().String("attribute-update-mode", "", "The method the specified attributes will update the existing attributes on the vehicle.")
		iotfleetwise_updateVehicleCmd.Flags().String("attributes", "", "Static information about a vehicle in a key-value pair.")
		iotfleetwise_updateVehicleCmd.Flags().String("decoder-manifest-arn", "", "The ARN of the decoder manifest associated with this vehicle.")
		iotfleetwise_updateVehicleCmd.Flags().String("model-manifest-arn", "", "The ARN of a vehicle model (model manifest) associated with the vehicle.")
		iotfleetwise_updateVehicleCmd.Flags().String("state-templates-to-add", "", "Associate state templates with the vehicle.")
		iotfleetwise_updateVehicleCmd.Flags().String("state-templates-to-remove", "", "Remove state templates from the vehicle.")
		iotfleetwise_updateVehicleCmd.Flags().String("state-templates-to-update", "", "Change the `stateTemplateUpdateStrategy` of state templates already associated with the vehicle.")
		iotfleetwise_updateVehicleCmd.Flags().String("vehicle-name", "", "The unique ID of the vehicle to update.")
		iotfleetwise_updateVehicleCmd.MarkFlagRequired("vehicle-name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_updateVehicleCmd)
}
