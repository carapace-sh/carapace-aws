package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_updateMonitorCmd = &cobra.Command{
	Use:   "update-monitor",
	Short: "Updates the `aggregationPeriod` for a monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_updateMonitorCmd).Standalone()

	networkmonitor_updateMonitorCmd.Flags().String("aggregation-period", "", "The aggregation time, in seconds, to change to.")
	networkmonitor_updateMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor to update.")
	networkmonitor_updateMonitorCmd.MarkFlagRequired("aggregation-period")
	networkmonitor_updateMonitorCmd.MarkFlagRequired("monitor-name")
	networkmonitorCmd.AddCommand(networkmonitor_updateMonitorCmd)
}
