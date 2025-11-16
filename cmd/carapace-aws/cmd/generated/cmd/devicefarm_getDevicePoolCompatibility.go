package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getDevicePoolCompatibilityCmd = &cobra.Command{
	Use:   "get-device-pool-compatibility",
	Short: "Gets information about compatibility with a device pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getDevicePoolCompatibilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getDevicePoolCompatibilityCmd).Standalone()

		devicefarm_getDevicePoolCompatibilityCmd.Flags().String("app-arn", "", "The ARN of the app that is associated with the specified device pool.")
		devicefarm_getDevicePoolCompatibilityCmd.Flags().String("configuration", "", "An object that contains information about the settings for a run.")
		devicefarm_getDevicePoolCompatibilityCmd.Flags().String("device-pool-arn", "", "The device pool's ARN.")
		devicefarm_getDevicePoolCompatibilityCmd.Flags().String("project-arn", "", "The ARN of the project for which you want to check device pool compatibility.")
		devicefarm_getDevicePoolCompatibilityCmd.Flags().String("test", "", "Information about the uploaded test to be run against the device pool.")
		devicefarm_getDevicePoolCompatibilityCmd.Flags().String("test-type", "", "The test type for the specified device pool.")
		devicefarm_getDevicePoolCompatibilityCmd.MarkFlagRequired("device-pool-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_getDevicePoolCompatibilityCmd)
}
