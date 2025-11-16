package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_deleteCoreDeviceCmd = &cobra.Command{
	Use:   "delete-core-device",
	Short: "Deletes a Greengrass core device, which is an IoT thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_deleteCoreDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_deleteCoreDeviceCmd).Standalone()

		greengrassv2_deleteCoreDeviceCmd.Flags().String("core-device-thing-name", "", "The name of the core device.")
		greengrassv2_deleteCoreDeviceCmd.MarkFlagRequired("core-device-thing-name")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_deleteCoreDeviceCmd)
}
