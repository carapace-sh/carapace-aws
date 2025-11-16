package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_updateDeviceInstanceCmd = &cobra.Command{
	Use:   "update-device-instance",
	Short: "Updates information about a private device instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_updateDeviceInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_updateDeviceInstanceCmd).Standalone()

		devicefarm_updateDeviceInstanceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the device instance.")
		devicefarm_updateDeviceInstanceCmd.Flags().String("labels", "", "An array of strings that you want to associate with the device instance.")
		devicefarm_updateDeviceInstanceCmd.Flags().String("profile-arn", "", "The ARN of the profile that you want to associate with the device instance.")
		devicefarm_updateDeviceInstanceCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_updateDeviceInstanceCmd)
}
