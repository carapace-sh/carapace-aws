package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_getMonitorCmd = &cobra.Command{
	Use:   "get-monitor",
	Short: "Gets information about a monitor in Network Flow Monitor based on a monitor name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_getMonitorCmd).Standalone()

	networkflowmonitor_getMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor.")
	networkflowmonitor_getMonitorCmd.MarkFlagRequired("monitor-name")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_getMonitorCmd)
}
