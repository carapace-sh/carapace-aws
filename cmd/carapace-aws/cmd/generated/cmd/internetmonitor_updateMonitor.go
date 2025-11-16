package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_updateMonitorCmd = &cobra.Command{
	Use:   "update-monitor",
	Short: "Updates a monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_updateMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitor_updateMonitorCmd).Standalone()

		internetmonitor_updateMonitorCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters that you specify to make an idempotent API request.")
		internetmonitor_updateMonitorCmd.Flags().String("health-events-config", "", "The list of health score thresholds.")
		internetmonitor_updateMonitorCmd.Flags().String("internet-measurements-log-delivery", "", "Publish internet measurements for Internet Monitor to another location, such as an Amazon S3 bucket.")
		internetmonitor_updateMonitorCmd.Flags().String("max-city-networks-to-monitor", "", "The maximum number of city-networks to monitor for your application.")
		internetmonitor_updateMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor.")
		internetmonitor_updateMonitorCmd.Flags().String("resources-to-add", "", "The resources to include in a monitor, which you provide as a set of Amazon Resource Names (ARNs).")
		internetmonitor_updateMonitorCmd.Flags().String("resources-to-remove", "", "The resources to remove from a monitor, which you provide as a set of Amazon Resource Names (ARNs).")
		internetmonitor_updateMonitorCmd.Flags().String("status", "", "The status for a monitor.")
		internetmonitor_updateMonitorCmd.Flags().String("traffic-percentage-to-monitor", "", "The percentage of the internet-facing traffic for your application that you want to monitor with this monitor.")
		internetmonitor_updateMonitorCmd.MarkFlagRequired("monitor-name")
	})
	internetmonitorCmd.AddCommand(internetmonitor_updateMonitorCmd)
}
