package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_describeDeviceCmd = &cobra.Command{
	Use:   "describe-device",
	Short: "Checks device-specific information, such as the device type, software version, IP addresses, and lock status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_describeDeviceCmd).Standalone()

	snowDeviceManagement_describeDeviceCmd.Flags().String("managed-device-id", "", "The ID of the device that you are checking the information of.")
	snowDeviceManagement_describeDeviceCmd.MarkFlagRequired("managed-device-id")
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_describeDeviceCmd)
}
