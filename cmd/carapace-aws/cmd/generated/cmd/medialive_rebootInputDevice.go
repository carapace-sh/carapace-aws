package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_rebootInputDeviceCmd = &cobra.Command{
	Use:   "reboot-input-device",
	Short: "Send a reboot command to the specified input device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_rebootInputDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_rebootInputDeviceCmd).Standalone()

		medialive_rebootInputDeviceCmd.Flags().String("force", "", "Force a reboot of an input device.")
		medialive_rebootInputDeviceCmd.Flags().String("input-device-id", "", "The unique ID of the input device to reboot.")
		medialive_rebootInputDeviceCmd.MarkFlagRequired("input-device-id")
	})
	medialiveCmd.AddCommand(medialive_rebootInputDeviceCmd)
}
