package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_deleteMonitorCmd = &cobra.Command{
	Use:   "delete-monitor",
	Short: "Deletes a monitor in Network Flow Monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_deleteMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_deleteMonitorCmd).Standalone()

		networkflowmonitor_deleteMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor to delete.")
		networkflowmonitor_deleteMonitorCmd.MarkFlagRequired("monitor-name")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_deleteMonitorCmd)
}
