package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deregisterDevicesCmd = &cobra.Command{
	Use:   "deregister-devices",
	Short: "Deregisters the specified devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deregisterDevicesCmd).Standalone()

	sagemaker_deregisterDevicesCmd.Flags().String("device-fleet-name", "", "The name of the fleet the devices belong to.")
	sagemaker_deregisterDevicesCmd.Flags().String("device-names", "", "The unique IDs of the devices.")
	sagemaker_deregisterDevicesCmd.MarkFlagRequired("device-fleet-name")
	sagemaker_deregisterDevicesCmd.MarkFlagRequired("device-names")
	sagemakerCmd.AddCommand(sagemaker_deregisterDevicesCmd)
}
