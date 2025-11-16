package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateDeviceFleetCmd = &cobra.Command{
	Use:   "update-device-fleet",
	Short: "Updates a fleet of devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateDeviceFleetCmd).Standalone()

	sagemaker_updateDeviceFleetCmd.Flags().String("description", "", "Description of the fleet.")
	sagemaker_updateDeviceFleetCmd.Flags().String("device-fleet-name", "", "The name of the fleet.")
	sagemaker_updateDeviceFleetCmd.Flags().String("enable-iot-role-alias", "", "Whether to create an Amazon Web Services IoT Role Alias during device fleet creation.")
	sagemaker_updateDeviceFleetCmd.Flags().String("output-config", "", "Output configuration for storing sample data collected by the fleet.")
	sagemaker_updateDeviceFleetCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the device.")
	sagemaker_updateDeviceFleetCmd.MarkFlagRequired("device-fleet-name")
	sagemaker_updateDeviceFleetCmd.MarkFlagRequired("output-config")
	sagemakerCmd.AddCommand(sagemaker_updateDeviceFleetCmd)
}
