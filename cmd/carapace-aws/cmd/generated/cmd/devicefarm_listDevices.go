package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Gets information about unique device types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listDevicesCmd).Standalone()

		devicefarm_listDevicesCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the project.")
		devicefarm_listDevicesCmd.Flags().String("filters", "", "Used to select a set of devices.")
		devicefarm_listDevicesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	})
	devicefarmCmd.AddCommand(devicefarm_listDevicesCmd)
}
