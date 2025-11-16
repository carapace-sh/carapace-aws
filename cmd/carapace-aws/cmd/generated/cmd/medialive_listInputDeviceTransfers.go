package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listInputDeviceTransfersCmd = &cobra.Command{
	Use:   "list-input-device-transfers",
	Short: "List input devices that are currently being transferred.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listInputDeviceTransfersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listInputDeviceTransfersCmd).Standalone()

		medialive_listInputDeviceTransfersCmd.Flags().String("max-results", "", "")
		medialive_listInputDeviceTransfersCmd.Flags().String("next-token", "", "")
		medialive_listInputDeviceTransfersCmd.Flags().String("transfer-type", "", "")
		medialive_listInputDeviceTransfersCmd.MarkFlagRequired("transfer-type")
	})
	medialiveCmd.AddCommand(medialive_listInputDeviceTransfersCmd)
}
