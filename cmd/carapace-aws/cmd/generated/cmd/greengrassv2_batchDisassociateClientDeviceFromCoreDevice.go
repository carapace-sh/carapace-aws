package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_batchDisassociateClientDeviceFromCoreDeviceCmd = &cobra.Command{
	Use:   "batch-disassociate-client-device-from-core-device",
	Short: "Disassociates a list of client devices from a core device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_batchDisassociateClientDeviceFromCoreDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_batchDisassociateClientDeviceFromCoreDeviceCmd).Standalone()

		greengrassv2_batchDisassociateClientDeviceFromCoreDeviceCmd.Flags().String("core-device-thing-name", "", "The name of the core device.")
		greengrassv2_batchDisassociateClientDeviceFromCoreDeviceCmd.Flags().String("entries", "", "The list of client devices to disassociate.")
		greengrassv2_batchDisassociateClientDeviceFromCoreDeviceCmd.MarkFlagRequired("core-device-thing-name")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_batchDisassociateClientDeviceFromCoreDeviceCmd)
}
