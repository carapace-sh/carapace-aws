package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateVtldeviceTypeCmd = &cobra.Command{
	Use:   "update-vtldevice-type",
	Short: "Updates the type of medium changer in a tape gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateVtldeviceTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_updateVtldeviceTypeCmd).Standalone()

		storagegateway_updateVtldeviceTypeCmd.Flags().String("device-type", "", "The type of medium changer you want to select.")
		storagegateway_updateVtldeviceTypeCmd.Flags().String("vtldevice-arn", "", "The Amazon Resource Name (ARN) of the medium changer you want to select.")
		storagegateway_updateVtldeviceTypeCmd.MarkFlagRequired("device-type")
		storagegateway_updateVtldeviceTypeCmd.MarkFlagRequired("vtldevice-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_updateVtldeviceTypeCmd)
}
