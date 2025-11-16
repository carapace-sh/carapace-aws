package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createDevicePoolCmd = &cobra.Command{
	Use:   "create-device-pool",
	Short: "Creates a device pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createDevicePoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_createDevicePoolCmd).Standalone()

		devicefarm_createDevicePoolCmd.Flags().String("description", "", "The device pool's description.")
		devicefarm_createDevicePoolCmd.Flags().String("max-devices", "", "The number of devices that Device Farm can add to your device pool.")
		devicefarm_createDevicePoolCmd.Flags().String("name", "", "The device pool's name.")
		devicefarm_createDevicePoolCmd.Flags().String("project-arn", "", "The ARN of the project for the device pool.")
		devicefarm_createDevicePoolCmd.Flags().String("rules", "", "The device pool's rules.")
		devicefarm_createDevicePoolCmd.MarkFlagRequired("name")
		devicefarm_createDevicePoolCmd.MarkFlagRequired("project-arn")
		devicefarm_createDevicePoolCmd.MarkFlagRequired("rules")
	})
	devicefarmCmd.AddCommand(devicefarm_createDevicePoolCmd)
}
