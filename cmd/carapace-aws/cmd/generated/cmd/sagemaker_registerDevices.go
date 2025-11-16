package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_registerDevicesCmd = &cobra.Command{
	Use:   "register-devices",
	Short: "Register devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_registerDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_registerDevicesCmd).Standalone()

		sagemaker_registerDevicesCmd.Flags().String("device-fleet-name", "", "The name of the fleet.")
		sagemaker_registerDevicesCmd.Flags().String("devices", "", "A list of devices to register with SageMaker Edge Manager.")
		sagemaker_registerDevicesCmd.Flags().String("tags", "", "The tags associated with devices.")
		sagemaker_registerDevicesCmd.MarkFlagRequired("device-fleet-name")
		sagemaker_registerDevicesCmd.MarkFlagRequired("devices")
	})
	sagemakerCmd.AddCommand(sagemaker_registerDevicesCmd)
}
