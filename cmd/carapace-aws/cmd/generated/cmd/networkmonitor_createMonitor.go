package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_createMonitorCmd = &cobra.Command{
	Use:   "create-monitor",
	Short: "Creates a monitor between a source subnet and destination IP address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_createMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmonitor_createMonitorCmd).Standalone()

		networkmonitor_createMonitorCmd.Flags().String("aggregation-period", "", "The time, in seconds, that metrics are aggregated and sent to Amazon CloudWatch.")
		networkmonitor_createMonitorCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure the idempotency of the request.")
		networkmonitor_createMonitorCmd.Flags().String("monitor-name", "", "The name identifying the monitor.")
		networkmonitor_createMonitorCmd.Flags().String("probes", "", "Displays a list of all of the probes created for a monitor.")
		networkmonitor_createMonitorCmd.Flags().String("tags", "", "The list of key-value pairs created and assigned to the monitor.")
		networkmonitor_createMonitorCmd.MarkFlagRequired("monitor-name")
	})
	networkmonitorCmd.AddCommand(networkmonitor_createMonitorCmd)
}
