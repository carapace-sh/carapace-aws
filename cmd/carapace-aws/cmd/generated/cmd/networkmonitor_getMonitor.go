package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_getMonitorCmd = &cobra.Command{
	Use:   "get-monitor",
	Short: "Returns details about a specific monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_getMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmonitor_getMonitorCmd).Standalone()

		networkmonitor_getMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor that details are returned for.")
		networkmonitor_getMonitorCmd.MarkFlagRequired("monitor-name")
	})
	networkmonitorCmd.AddCommand(networkmonitor_getMonitorCmd)
}
