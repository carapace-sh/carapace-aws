package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_transferInputDeviceCmd = &cobra.Command{
	Use:   "transfer-input-device",
	Short: "Start an input device transfer to another AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_transferInputDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_transferInputDeviceCmd).Standalone()

		medialive_transferInputDeviceCmd.Flags().String("input-device-id", "", "The unique ID of this input device.")
		medialive_transferInputDeviceCmd.Flags().String("target-customer-id", "", "The AWS account ID (12 digits) for the recipient of the device transfer.")
		medialive_transferInputDeviceCmd.Flags().String("target-region", "", "The target AWS region to transfer the device.")
		medialive_transferInputDeviceCmd.Flags().String("transfer-message", "", "An optional message for the recipient.")
		medialive_transferInputDeviceCmd.MarkFlagRequired("input-device-id")
	})
	medialiveCmd.AddCommand(medialive_transferInputDeviceCmd)
}
