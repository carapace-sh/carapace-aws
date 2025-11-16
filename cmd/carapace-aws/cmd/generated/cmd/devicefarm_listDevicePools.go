package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listDevicePoolsCmd = &cobra.Command{
	Use:   "list-device-pools",
	Short: "Gets information about device pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listDevicePoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listDevicePoolsCmd).Standalone()

		devicefarm_listDevicePoolsCmd.Flags().String("arn", "", "The project ARN.")
		devicefarm_listDevicePoolsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
		devicefarm_listDevicePoolsCmd.Flags().String("type", "", "The device pools' type.")
		devicefarm_listDevicePoolsCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listDevicePoolsCmd)
}
