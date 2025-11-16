package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_acceptInputDeviceTransferCmd = &cobra.Command{
	Use:   "accept-input-device-transfer",
	Short: "Accept an incoming input device transfer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_acceptInputDeviceTransferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_acceptInputDeviceTransferCmd).Standalone()

		medialive_acceptInputDeviceTransferCmd.Flags().String("input-device-id", "", "The unique ID of the input device to accept.")
		medialive_acceptInputDeviceTransferCmd.MarkFlagRequired("input-device-id")
	})
	medialiveCmd.AddCommand(medialive_acceptInputDeviceTransferCmd)
}
