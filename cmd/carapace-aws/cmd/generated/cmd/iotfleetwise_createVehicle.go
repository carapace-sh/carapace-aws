package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_createVehicleCmd = &cobra.Command{
	Use:   "create-vehicle",
	Short: "Creates a vehicle, which is an instance of a vehicle model (model manifest).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_createVehicleCmd).Standalone()

	iotfleetwise_createVehicleCmd.Flags().String("association-behavior", "", "An option to create a new Amazon Web Services IoT thing when creating a vehicle, or to validate an existing Amazon Web Services IoT thing as a vehicle.")
	iotfleetwise_createVehicleCmd.Flags().String("attributes", "", "Static information about a vehicle in a key-value pair.")
	iotfleetwise_createVehicleCmd.Flags().String("decoder-manifest-arn", "", "The ARN of a decoder manifest.")
	iotfleetwise_createVehicleCmd.Flags().String("model-manifest-arn", "", "The Amazon Resource Name ARN of a vehicle model.")
	iotfleetwise_createVehicleCmd.Flags().String("state-templates", "", "Associate state templates with the vehicle.")
	iotfleetwise_createVehicleCmd.Flags().String("tags", "", "Metadata that can be used to manage the vehicle.")
	iotfleetwise_createVehicleCmd.Flags().String("vehicle-name", "", "The unique ID of the vehicle to create.")
	iotfleetwise_createVehicleCmd.MarkFlagRequired("decoder-manifest-arn")
	iotfleetwise_createVehicleCmd.MarkFlagRequired("model-manifest-arn")
	iotfleetwise_createVehicleCmd.MarkFlagRequired("vehicle-name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_createVehicleCmd)
}
