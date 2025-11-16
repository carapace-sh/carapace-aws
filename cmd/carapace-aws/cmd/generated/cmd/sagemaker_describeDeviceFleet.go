package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeDeviceFleetCmd = &cobra.Command{
	Use:   "describe-device-fleet",
	Short: "A description of the fleet the device belongs to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeDeviceFleetCmd).Standalone()

	sagemaker_describeDeviceFleetCmd.Flags().String("device-fleet-name", "", "The name of the fleet.")
	sagemaker_describeDeviceFleetCmd.MarkFlagRequired("device-fleet-name")
	sagemakerCmd.AddCommand(sagemaker_describeDeviceFleetCmd)
}
