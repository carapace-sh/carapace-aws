package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getDeviceInstanceCmd = &cobra.Command{
	Use:   "get-device-instance",
	Short: "Returns information about a device instance that belongs to a private device fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getDeviceInstanceCmd).Standalone()

	devicefarm_getDeviceInstanceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the instance you're requesting information about.")
	devicefarm_getDeviceInstanceCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_getDeviceInstanceCmd)
}
