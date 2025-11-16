package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_createMonitorCmd = &cobra.Command{
	Use:   "create-monitor",
	Short: "Creates a monitor in Amazon CloudWatch Internet Monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_createMonitorCmd).Standalone()

	internetmonitor_createMonitorCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters that you specify to make an idempotent API request.")
	internetmonitor_createMonitorCmd.Flags().String("health-events-config", "", "Defines the threshold percentages and other configuration information for when Amazon CloudWatch Internet Monitor creates a health event.")
	internetmonitor_createMonitorCmd.Flags().String("internet-measurements-log-delivery", "", "Publish internet measurements for Internet Monitor to an Amazon S3 bucket in addition to CloudWatch Logs.")
	internetmonitor_createMonitorCmd.Flags().String("max-city-networks-to-monitor", "", "The maximum number of city-networks to monitor for your resources.")
	internetmonitor_createMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor.")
	internetmonitor_createMonitorCmd.Flags().String("resources", "", "The resources to include in a monitor, which you provide as a set of Amazon Resource Names (ARNs).")
	internetmonitor_createMonitorCmd.Flags().String("tags", "", "The tags for a monitor.")
	internetmonitor_createMonitorCmd.Flags().String("traffic-percentage-to-monitor", "", "The percentage of the internet-facing traffic for your application that you want to monitor with this monitor.")
	internetmonitor_createMonitorCmd.MarkFlagRequired("monitor-name")
	internetmonitorCmd.AddCommand(internetmonitor_createMonitorCmd)
}
