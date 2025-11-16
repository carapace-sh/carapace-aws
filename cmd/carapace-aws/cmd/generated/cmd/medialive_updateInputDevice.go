package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateInputDeviceCmd = &cobra.Command{
	Use:   "update-input-device",
	Short: "Updates the parameters for the input device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateInputDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_updateInputDeviceCmd).Standalone()

		medialive_updateInputDeviceCmd.Flags().String("availability-zone", "", "The Availability Zone you want associated with this input device.")
		medialive_updateInputDeviceCmd.Flags().String("hd-device-settings", "", "The settings that you want to apply to the HD input device.")
		medialive_updateInputDeviceCmd.Flags().String("input-device-id", "", "The unique ID of the input device.")
		medialive_updateInputDeviceCmd.Flags().String("name", "", "The name that you assigned to this input device (not the unique ID).")
		medialive_updateInputDeviceCmd.Flags().String("uhd-device-settings", "", "The settings that you want to apply to the UHD input device.")
		medialive_updateInputDeviceCmd.MarkFlagRequired("input-device-id")
	})
	medialiveCmd.AddCommand(medialive_updateInputDeviceCmd)
}
