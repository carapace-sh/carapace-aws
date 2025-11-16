package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_updateDevicePoolCmd = &cobra.Command{
	Use:   "update-device-pool",
	Short: "Modifies the name, description, and rules in a device pool given the attributes and the pool ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_updateDevicePoolCmd).Standalone()

	devicefarm_updateDevicePoolCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Device Farm device pool to update.")
	devicefarm_updateDevicePoolCmd.Flags().Bool("clear-max-devices", false, "Sets whether the `maxDevices` parameter applies to your device pool.")
	devicefarm_updateDevicePoolCmd.Flags().String("description", "", "A description of the device pool to update.")
	devicefarm_updateDevicePoolCmd.Flags().String("max-devices", "", "The number of devices that Device Farm can add to your device pool.")
	devicefarm_updateDevicePoolCmd.Flags().String("name", "", "A string that represents the name of the device pool to update.")
	devicefarm_updateDevicePoolCmd.Flags().Bool("no-clear-max-devices", false, "Sets whether the `maxDevices` parameter applies to your device pool.")
	devicefarm_updateDevicePoolCmd.Flags().String("rules", "", "Represents the rules to modify for the device pool.")
	devicefarm_updateDevicePoolCmd.MarkFlagRequired("arn")
	devicefarm_updateDevicePoolCmd.Flag("no-clear-max-devices").Hidden = true
	devicefarmCmd.AddCommand(devicefarm_updateDevicePoolCmd)
}
