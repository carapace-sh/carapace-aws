package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_stopInputDeviceCmd = &cobra.Command{
	Use:   "stop-input-device",
	Short: "Stop an input device that is attached to a MediaConnect flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_stopInputDeviceCmd).Standalone()

	medialive_stopInputDeviceCmd.Flags().String("input-device-id", "", "The unique ID of the input device to stop.")
	medialive_stopInputDeviceCmd.MarkFlagRequired("input-device-id")
	medialiveCmd.AddCommand(medialive_stopInputDeviceCmd)
}
