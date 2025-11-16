package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_cancelInputDeviceTransferCmd = &cobra.Command{
	Use:   "cancel-input-device-transfer",
	Short: "Cancel an input device transfer that you have requested.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_cancelInputDeviceTransferCmd).Standalone()

	medialive_cancelInputDeviceTransferCmd.Flags().String("input-device-id", "", "The unique ID of the input device to cancel.")
	medialive_cancelInputDeviceTransferCmd.MarkFlagRequired("input-device-id")
	medialiveCmd.AddCommand(medialive_cancelInputDeviceTransferCmd)
}
