package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateDevicesCmd = &cobra.Command{
	Use:   "update-devices",
	Short: "Updates one or more devices in a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateDevicesCmd).Standalone()

		sagemaker_updateDevicesCmd.Flags().String("device-fleet-name", "", "The name of the fleet the devices belong to.")
		sagemaker_updateDevicesCmd.Flags().String("devices", "", "List of devices to register with Edge Manager agent.")
		sagemaker_updateDevicesCmd.MarkFlagRequired("device-fleet-name")
		sagemaker_updateDevicesCmd.MarkFlagRequired("devices")
	})
	sagemakerCmd.AddCommand(sagemaker_updateDevicesCmd)
}
