package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describeDeviceCmd = &cobra.Command{
	Use:   "describe-device",
	Short: "Returns information about a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describeDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_describeDeviceCmd).Standalone()

		panorama_describeDeviceCmd.Flags().String("device-id", "", "The device's ID.")
		panorama_describeDeviceCmd.MarkFlagRequired("device-id")
	})
	panoramaCmd.AddCommand(panorama_describeDeviceCmd)
}
