package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_rejectInputDeviceTransferCmd = &cobra.Command{
	Use:   "reject-input-device-transfer",
	Short: "Reject the transfer of the specified input device to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_rejectInputDeviceTransferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_rejectInputDeviceTransferCmd).Standalone()

		medialive_rejectInputDeviceTransferCmd.Flags().String("input-device-id", "", "The unique ID of the input device to reject.")
		medialive_rejectInputDeviceTransferCmd.MarkFlagRequired("input-device-id")
	})
	medialiveCmd.AddCommand(medialive_rejectInputDeviceTransferCmd)
}
