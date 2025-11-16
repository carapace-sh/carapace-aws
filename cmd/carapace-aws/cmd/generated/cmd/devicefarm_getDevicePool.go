package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getDevicePoolCmd = &cobra.Command{
	Use:   "get-device-pool",
	Short: "Gets information about a device pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getDevicePoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getDevicePoolCmd).Standalone()

		devicefarm_getDevicePoolCmd.Flags().String("arn", "", "The device pool's ARN.")
		devicefarm_getDevicePoolCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_getDevicePoolCmd)
}
