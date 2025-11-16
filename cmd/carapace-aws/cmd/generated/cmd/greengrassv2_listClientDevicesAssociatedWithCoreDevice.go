package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_listClientDevicesAssociatedWithCoreDeviceCmd = &cobra.Command{
	Use:   "list-client-devices-associated-with-core-device",
	Short: "Retrieves a paginated list of client devices that are associated with a core device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_listClientDevicesAssociatedWithCoreDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_listClientDevicesAssociatedWithCoreDeviceCmd).Standalone()

		greengrassv2_listClientDevicesAssociatedWithCoreDeviceCmd.Flags().String("core-device-thing-name", "", "The name of the core device.")
		greengrassv2_listClientDevicesAssociatedWithCoreDeviceCmd.Flags().String("max-results", "", "The maximum number of results to be returned per paginated request.")
		greengrassv2_listClientDevicesAssociatedWithCoreDeviceCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		greengrassv2_listClientDevicesAssociatedWithCoreDeviceCmd.MarkFlagRequired("core-device-thing-name")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_listClientDevicesAssociatedWithCoreDeviceCmd)
}
