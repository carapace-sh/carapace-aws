package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Returns a list of all devices on your Amazon Web Services account that have Amazon Web Services Snow Device Management enabled in the Amazon Web Services Region where the command is run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_listDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagement_listDevicesCmd).Standalone()

		snowDeviceManagement_listDevicesCmd.Flags().String("job-id", "", "The ID of the job used to order the device.")
		snowDeviceManagement_listDevicesCmd.Flags().String("max-results", "", "The maximum number of devices to list per page.")
		snowDeviceManagement_listDevicesCmd.Flags().String("next-token", "", "A pagination token to continue to the next page of results.")
	})
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_listDevicesCmd)
}
