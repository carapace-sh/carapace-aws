package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createDeviceFleetCmd = &cobra.Command{
	Use:   "create-device-fleet",
	Short: "Creates a device fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createDeviceFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createDeviceFleetCmd).Standalone()

		sagemaker_createDeviceFleetCmd.Flags().String("description", "", "A description of the fleet.")
		sagemaker_createDeviceFleetCmd.Flags().String("device-fleet-name", "", "The name of the fleet that the device belongs to.")
		sagemaker_createDeviceFleetCmd.Flags().String("enable-iot-role-alias", "", "Whether to create an Amazon Web Services IoT Role Alias during device fleet creation.")
		sagemaker_createDeviceFleetCmd.Flags().String("output-config", "", "The output configuration for storing sample data collected by the fleet.")
		sagemaker_createDeviceFleetCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) that has access to Amazon Web Services Internet of Things (IoT).")
		sagemaker_createDeviceFleetCmd.Flags().String("tags", "", "Creates tags for the specified fleet.")
		sagemaker_createDeviceFleetCmd.MarkFlagRequired("device-fleet-name")
		sagemaker_createDeviceFleetCmd.MarkFlagRequired("output-config")
	})
	sagemakerCmd.AddCommand(sagemaker_createDeviceFleetCmd)
}
