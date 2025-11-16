package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listWirelessDeviceImportTasksCmd = &cobra.Command{
	Use:   "list-wireless-device-import-tasks",
	Short: "List of import tasks and summary information of onboarding status of devices in each import task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listWirelessDeviceImportTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listWirelessDeviceImportTasksCmd).Standalone()

		iotwireless_listWirelessDeviceImportTasksCmd.Flags().String("max-results", "", "")
		iotwireless_listWirelessDeviceImportTasksCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise `null` to receive the first set of results.")
	})
	iotwirelessCmd.AddCommand(iotwireless_listWirelessDeviceImportTasksCmd)
}
