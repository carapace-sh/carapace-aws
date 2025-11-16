package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_getCoreDeviceCmd = &cobra.Command{
	Use:   "get-core-device",
	Short: "Retrieves metadata for a Greengrass core device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_getCoreDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_getCoreDeviceCmd).Standalone()

		greengrassv2_getCoreDeviceCmd.Flags().String("core-device-thing-name", "", "The name of the core device.")
		greengrassv2_getCoreDeviceCmd.MarkFlagRequired("core-device-thing-name")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_getCoreDeviceCmd)
}
