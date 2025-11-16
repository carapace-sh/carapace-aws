package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteDevicePoolCmd = &cobra.Command{
	Use:   "delete-device-pool",
	Short: "Deletes a device pool given the pool ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteDevicePoolCmd).Standalone()

	devicefarm_deleteDevicePoolCmd.Flags().String("arn", "", "Represents the Amazon Resource Name (ARN) of the Device Farm device pool to delete.")
	devicefarm_deleteDevicePoolCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_deleteDevicePoolCmd)
}
