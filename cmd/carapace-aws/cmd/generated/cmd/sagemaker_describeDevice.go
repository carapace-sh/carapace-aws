package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeDeviceCmd = &cobra.Command{
	Use:   "describe-device",
	Short: "Describes the device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeDeviceCmd).Standalone()

	sagemaker_describeDeviceCmd.Flags().String("device-fleet-name", "", "The name of the fleet the devices belong to.")
	sagemaker_describeDeviceCmd.Flags().String("device-name", "", "The unique ID of the device.")
	sagemaker_describeDeviceCmd.Flags().String("next-token", "", "Next token of device description.")
	sagemaker_describeDeviceCmd.MarkFlagRequired("device-fleet-name")
	sagemaker_describeDeviceCmd.MarkFlagRequired("device-name")
	sagemakerCmd.AddCommand(sagemaker_describeDeviceCmd)
}
