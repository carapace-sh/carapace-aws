package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_batchAssociateClientDeviceWithCoreDeviceCmd = &cobra.Command{
	Use:   "batch-associate-client-device-with-core-device",
	Short: "Associates a list of client devices with a core device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_batchAssociateClientDeviceWithCoreDeviceCmd).Standalone()

	greengrassv2_batchAssociateClientDeviceWithCoreDeviceCmd.Flags().String("core-device-thing-name", "", "The name of the core device.")
	greengrassv2_batchAssociateClientDeviceWithCoreDeviceCmd.Flags().String("entries", "", "The list of client devices to associate.")
	greengrassv2_batchAssociateClientDeviceWithCoreDeviceCmd.MarkFlagRequired("core-device-thing-name")
	greengrassv2Cmd.AddCommand(greengrassv2_batchAssociateClientDeviceWithCoreDeviceCmd)
}
