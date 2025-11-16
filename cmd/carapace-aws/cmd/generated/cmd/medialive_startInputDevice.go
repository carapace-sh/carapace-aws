package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_startInputDeviceCmd = &cobra.Command{
	Use:   "start-input-device",
	Short: "Start an input device that is attached to a MediaConnect flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_startInputDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_startInputDeviceCmd).Standalone()

		medialive_startInputDeviceCmd.Flags().String("input-device-id", "", "The unique ID of the input device to start.")
		medialive_startInputDeviceCmd.MarkFlagRequired("input-device-id")
	})
	medialiveCmd.AddCommand(medialive_startInputDeviceCmd)
}
